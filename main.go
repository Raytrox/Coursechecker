package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	r := Router()

	fmt.Println("Starting server on port 8010...")
	log.Fatal(http.ListenAndServe(":8010", r))
}
