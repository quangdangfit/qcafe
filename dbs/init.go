package dbs

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"qcafe/config"
	"time"
)

var MongoSession *mgo.Session

func init() {
	info := &mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%d", config.Config.MongoDB.Host, config.Config.MongoDB.Port)},
		Timeout:  60 * time.Second,
		Database: config.Config.MongoDB.AuthDatabase,
		Username: config.Config.MongoDB.User,
		Password: config.Config.MongoDB.Password,
	}
	mgoSession, err := mgo.DialWithInfo(info)

	if err != nil {
		log.Fatal("Cannot connect to Database")
	}

	MongoSession = mgoSession
}
