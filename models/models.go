package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Unit is the one monitored course, contains the course and user information
type Unit struct {
	ID primitive.ObjectID		`json:"id"`
	Dept 		string			`json:"dept"`
	Number 		string			`json:"number"`
	Section 	string			`json:"section"`
	Receiver 	string			`json:"receiver"`
	Status		string			`json:"status"`
}
