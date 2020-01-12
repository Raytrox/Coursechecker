package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"github.com/magiconair/properties"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"
	"./models"
)

// DB connection string
const connection = "mongodb+srv://admin:960925@learning-0wti5.mongodb.net/test?retryWrites=true&w=majority"

// DB name
const dbName = "coursechecker"

// Collection name
const collName = "monitoring"

// collection instance
var collection *mongo.Collection

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(connection)

	// connect to MangoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to MangoDB!")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection created!")
}

// GetAllUnits from database
func GetAllUnits(w http.ResponseWriter, r *http.Request) {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Get request received, begin to process request \n")
	file := "cc.conf"
	p := properties.MustLoadFile(file, properties.UTF8)
	pass := p.MustGetString("SENDERPASS")
	sender := p.MustGetString("SENDER")

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var course models.Unit

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
	}

	addUnit(course)

	for {
		if getNumberofSeats(course) != 0 {
			notifyUser(course, pass, sender)
			break
		}
		fmt.Printf("Currently there is no seat available for the course\n")
		time.Sleep(10000 * time.Millisecond)
	}
	// TODO: use multithread to handle multiple request and add response
	// response could either be monitoring started with 200 or bad data input with 400 
}

// getNumberofSeats takes dept, course number and section number
// returns the current general seats remaining for the section
func getNumberofSeats(c models.Unit) int {
	temp := 0
	url := "https://courses.students.ubc.ca/cs/courseschedule?pname=subjarea&tname=subj-section&" +
		"dept=" + c.Dept + "&course=" + c.Number + "&section=" + c.Section
	fmt.Println("the url requested is: " + url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("status code error: %s %d\n", res.Status, res.StatusCode)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("table.\\'table").Each(func(index int, item *goquery.Selection) {
		info := item.Text()
		seats := strings.Split(info, "\n")
		temp, _ = strconv.Atoi(strings.SplitAfter(seats[4], "General Seats Remaining:")[1])
	})
	return temp
}

func notifyUser(c models.Unit, pass string, sender string) {
	from := sender

	msg := "From: " + from + "\n" +
		"To: " + c.Receiver + "\n" +
		"Subject: Hello there from Course Checker \n\n" +
		c.Dept + c.Number + " " + c.Section +
		" is available now, please register ASAP."

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{c.Receiver}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("email sent, please restart the program if you want to monitor the registration again\n")
}

func addUnit(unit models.Unit) {
	result, err := collection.InsertOne(context.Background(), unit)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one course into the database ", result.InsertedID)
}


