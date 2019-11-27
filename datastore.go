package main

import (
	"gopkg.in/mgo.v2"
)

// CreateSession creates the main session to our MongoDB instane
func CreateSession(host string) (*mgo.Session, error) {
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
