package main

import (
    // "bytes"
    // "encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	"net/smtp"
	"log"
	"os"
)

func main() {
    fmt.Println("Starting the application...")
    response, err := http.Get("https://courses.students.ubc.ca/cs/courseschedule?pname=subjarea&tname=subj-section&dept=GERM&course=433&section=905")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
	}
	notifyUser()
    // jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
    // jsonValue, _ := json.Marshal(jsonData)
    // response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
    // if err != nil {
    //     fmt.Printf("The HTTP request failed with error %s\n", err)
    // } else {
    //     data, _ := ioutil.ReadAll(response.Body)
    //     fmt.Println(string(data))
    // }
    // fmt.Println("Terminating the application...")
}

func notifyUser() {
	from := "ruisharp25@gmail.com"
	pass := os.Getenv("SENDERPASS");
	to := "yinsharp25@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there from Course Checker \n\n" +
		"Your course is available now, please register ASAP."

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	
	log.Print("sent, visit http://foobarbazz.mailinator.com")
}