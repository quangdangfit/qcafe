package dbs

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

var MongoSession *mgo.Session

func init() {
	info := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Timeout:  60 * time.Second,
		Database: "admin",
	}
	mgoSession, err := mgo.DialWithInfo(info)

	if err != nil {
		fmt.Println("Cannot connect to Mongo DB")
	}

	MongoSession = mgoSession
}
