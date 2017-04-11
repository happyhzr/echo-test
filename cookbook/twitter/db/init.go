package db

import (
	"github.com/insisthzr/echo-test/cookbook/twitter/conf"

	"gopkg.in/mgo.v2"
)

var (
	sess *mgo.Session
)

func init() {
	err := InitDB()
	if err != nil {
		panic(err)
	}
	err = ensureIndex()
	if err != nil {
		panic(err)
	}

}

// InitDB InitDB
func InitDB() error {
	var err error
	sess, err = mgo.Dial(conf.MONGODB_HOST)
	if err != nil {
		return err
	}
	sess.SetMode(mgo.Monotonic, true)
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
