package db

import (
	"os"

	"gopkg.in/mgo.v2"
)

var (
	sess *mgo.Session
)

// InitDB InitDB
func InitDB() error {
	var mongoURI string
	if os.Getenv("MONGODB") == "" {
		mongoURI = "localhost"
	} else {
		mongoURI = os.Getenv("MONGODB")
	}
	var err error
	sess, err = mgo.Dial(mongoURI)
	if err != nil {
		return err
	}
	sess.SetMode(mgo.Monotonic, true)

	err = EnsureIndex()
	if err != nil {
		return err
	}

	return nil
}

// CheckStatus check db session status.
func CheckStatus() error {
	return sess.Ping()
}

// NewDBSession NewDBSession
func NewDBSession() *mgo.Session {
	return sess.Clone()
}
