package main

import (
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

// the mongoDB Server adress
const mongoServerAddress = "mongodb://3baule:bathesis16@ds139197.mlab.com:39197/3baule_database"

// creates a global session that connects with my mLab database
var session, err = mgo.Dial(mongoServerAddress)

// Main Function that sets up a server with the Functions that it can handle
func main() {

	// if an error occurs, then panic
	if err != nil {
		panic(err)
	}

	// closes the session after it is no longer needed
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// creates a new Router
	router := NewRouter()

	// connect to the server address
	log.Fatal(http.ListenAndServe(":8080", router))
}
