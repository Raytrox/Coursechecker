package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strconv"
	"strings"
	"time"
	// "io/ioutil"
	"net/http"
	"net/smtp"
	"os"
)

type Course struct {
	dept string
	course string
	section string
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage:", os.Args[0], "DEPT", "COURSE", "SECTION")
		return
	}
	dept := strings.ToLower(os.Args[1])
	course := strings.ToLower(os.Args[2])
	section := strings.ToLower(os.Args[3])
	temp := Course{dept:dept, course:course, section:section}
	for {
		if getNumberofSeats(temp) != 0 {
			notifyUser(temp)
			break
		}
		fmt.Printf("Currently there is no seat available for the course\n")
		time.Sleep(10000 * time.Millisecond)
	}
}

// getNumberofSeats takes dept, course number and section number
// returns the current general seats remaining for the section
func getNumberofSeats(course Course) int {
	temp := 0
	url := "https://courses.students.ubc.ca/cs/courseschedule?pname=subjarea&tname=subj-section&" + 
	"dept=" + course.dept + "&course=" + course.course + "&section=" + course.section
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("status code error: %s %d\n", res.Status, res.StatusCode)
	}
	doc,err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("table.\\'table").Each(func(index int, item *goquery.Selection) {
		info := item.Text()
		seats := strings.Split(info, "\n")
		temp,_ = strconv.Atoi(strings.SplitAfter(seats[4], "General Seats Remaining:")[1])
	})
	return temp
}

func notifyUser(course Course) {
	from := "ruisharp25@gmail.com"
	pass := os.Getenv("SENDERPASS");
	to := "yinsharp25@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there from Course Checker \n\n" +
		course.dept + course.course + " " + course.section +
		" is available now, please register ASAP."

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	
	log.Print("email sent, please restart the program if you want to monitor the registration again\n")
}