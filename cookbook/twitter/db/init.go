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
	sess, err := mgo.Dial(mongoURI)
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

// EnsureIndex create index
func EnsureIndex() error {
	index := mgo.Index{Key: []string{"email"},
		Unique: true,
	}
	err := sess.Copy().DB("twitter").C("users").EnsureIndex(index)
	return err
}

// CheckStatus check db session status.
func CheckStatus() bool {
	return sess.Ping() == nil
}

// NewDBSession NewDBSession
func NewDBSession() *mgo.Session {
	return sess.Clone()
}
