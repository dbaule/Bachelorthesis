package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Index hnadles the "/" statement on the server by saying hello
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// CreateViewController creates a new viewcontroller in the viewcontroller collection
func CreateViewController(s *mgo.Session) http.HandlerFunc {

	// the returned HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {

		// copies the session to work with it
		session := s.Copy()

		// the viewcontroller that will be created
		var viewcontroller ViewController

		// reads the body of the request until a certain limit
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

		if err != nil {
			panic(err)
		}

		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		// closes the session after it is no longer needed
		defer session.Close()

		// unmarshals the body of the request when there were no previous errors
		if err := json.Unmarshal(body, &viewcontroller); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")

			w.WriteHeader(422) // unprocessible entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		// gets the viewcontroller collection and inserts the created ViewController
		viewcontrollers := session.DB("3baule_database").C("viewcontrollerss")

		// if a viewcontroller with the same name and buildversion already exists, write header code 208,
		// otherwise add it to the collection
		err = viewcontrollers.Insert(viewcontroller)
		if err != nil {

			if mgo.IsDup(err) {
				w.WriteHeader(http.StatusAlreadyReported)
			}
		} else {

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusCreated)

			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

	}

}

// GetViewController gets two ViewController with the same name, one having the highest
// buildversion currently available in the db and the other with the previous number
func GetViewController(s *mgo.Session) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// copies the session to work with it
		session := s.Copy()
		// closes the session after it is no longer needed
		defer session.Close()
		// the database
		db := session.DB("3baule_database")
		// the viewcontrollers collection
		vc := db.C("viewcontrollers")

		name := "ViewController"

		// gets the Name of the ViewController to compare
		vcName := r.URL.Query().Get("vc")

		if len(vcName) != 0 {
			name = vcName
		} else {
			fmt.Println("Please provide a ViewController name, e.g.: getViewController?vc=ViewController1")
			return
		}

		// orders the ViewController by build_version descending and saves the latest two to the vcs array
		vcs := make([]ViewController, 0)
		vc.Find(bson.M{"name": name}).Sort("-build_version").Limit(2).All(&vcs)

		//count, err := viewcontrollers.Count()
		if err != nil {
			panic(err)
		}

		// for _, vc := range vcs {
		// 	fmt.Println(vc)
		// }

		// when two ViewController were found they will get compared
		if len(vcs) > 1 {
			compareViewController(vcs[0], vcs[1])
		} else {
			fmt.Println("There was no ViewController with the name " + name)
		}
	}
}

// GetDifferences gets all differences that can be found in the app
func GetDifferences(s *mgo.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetDifferences is called")
		// copies the session
		session := s.Copy()
		//closes the session after it is no longer needed
		defer session.Close()

		// the database 3baule_database
		db := session.DB("3baule_database")
		// the viewcontrollers collection
		vc := db.C("viewcontrollerss")

		// gets the first ViewController after sorting for build_version descending
		var firstVC ViewController
		vc.Find(nil).Sort("-build_version").One(&firstVC)
		// get firstVC build_version for the current highest build_version
		highestBuildVersion := firstVC.BuildVersion
		fmt.Println("highest build " + highestBuildVersion)

		// convert the string to an int and substract by 1 and save the new number as a string to secondHighestNum
		i, err := strconv.Atoi(highestBuildVersion)
		secondHighestNum := ""
		if err == nil {
			secondHighestNum = strconv.Itoa(i - 1)
		}
		fmt.Println("second highest: " + secondHighestNum)

		// gets all ViewController with the highestBuildVersion
		vcs1 := make([]ViewController, 0)
		vc.Find(bson.M{"build_version": highestBuildVersion}).All(&vcs1)

		// gets all ViewController with the secondHighestNum as build_version
		vcs2 := make([]ViewController, 0)
		vc.Find(bson.M{"build_version": secondHighestNum}).All(&vcs2)

		// calls compareAllViewController with both arrays to compare the individual ViewController with each other
		compareAllViewController(vcs1, vcs2, w, r)
	}
}

// CreateEvent creates a new event in the event collection
func CreateEvent(s *mgo.Session) http.HandlerFunc {

	// the returned HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {

		// copies the session to work with it
		session := s.Copy()

		// the Event that will be created
		var event Event

		// reads the body of the request until a certain limit
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

		if err != nil {
			panic(err)
		}

		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		// closes the session after it is no longer needed
		defer session.Close()

		// unmarshals the body of the request when there were no previous errors
		if err := json.Unmarshal(body, &event); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")

			w.WriteHeader(422) // unprocessible entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		// gets the events collection and inserts the created Event
		events := session.DB("3baule_database").C("events")

		// if a event with the same properties exists the statuscode 208 is written and
		// the event will not be inserted into the collection
		err = events.Insert(event)
		if err != nil {

			if mgo.IsDup(err) {
				w.WriteHeader(http.StatusAlreadyReported)
			}
		} else {

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusCreated)

			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

	}

}
