package db

import (
	"github.com/insisthzr/echo-test/cookbook/twitter/conf"

	"gopkg.in/mgo.v2"
)

// EnsureIndex create index
func ensureIndex() error {
	sess := NewDBSession()
	defer sess.Close()

	err := sess.DB(conf.DB_NAME).C("users").EnsureIndexKey(mgo.Index{Key: []string{"email"},
		Unique: true,
	})
	if err != nil {
		return err
	}
	err = sess.DB(conf.DB_NAME).C("posts").EnsureIndexKey("to")
	if err != nil {
		return err
	}
	err = sess.DB(conf.DB_NAME).C("posts").EnsureIndexKey("from")
	if err != nil {
		return err
	}
	return nil
}
