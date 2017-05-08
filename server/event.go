package main

// Event is struct that holds a transition that occured in an app
//TODO: omitempty for properties that should have a value
type Event struct {
	Source      string `json:"source" bson:"source"`
	Name        string `json:"name" bson:"name"`
	Destination string `json:"destination" bson:"destination"`
}

// Events holds a collection of Event's
type Events []Event
