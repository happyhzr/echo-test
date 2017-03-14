package db

import (
	"github.com/insisthzr/echo-test/cookbook/twitter/conf"

	"gopkg.in/mgo.v2"
)

// EnsureIndex create index
func EnsureIndex() error {
	sess := NewDBSession()
	defer sess.Close()

	err := sess.DB(conf.DBName).C("users").EnsureIndex(mgo.Index{Key: []string{"email"},
		Unique: true,
	})
	if err != nil {
		return err
	}
	err = sess.DB(conf.DBName).C("posts").EnsureIndexKey("to")
	if err != nil {
		return err
	}
	err = sess.DB(conf.DBName).C("posts").EnsureIndexKey("from")
	if err != nil {
		return err
	}
	return nil
}
