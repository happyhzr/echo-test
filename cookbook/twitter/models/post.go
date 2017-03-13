package models

import (
	//"os/user"

	//"github.com/insisthzr/echo-test/cookbook/twitter/conf"
	//"github.com/insisthzr/echo-test/cookbook/twitter/db"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	To      string        `json:"to" bson:"to"`
	From    string        `json:"from" bson:"from"`
	Message string        `json:"message" bson:"message"`
}

/*
func AddPost(id bson.ObjectId,post *Post) error {
	sess := db.NewDBSession()
	defer sess.Close()

	err = sess.DB(conf.DBName).C("users").FindId(user.ID).One(user)
	if err != nil {
		return err
	}

	err = db.DB("twitter").C("posts").Insert(post)
	if err != nil {
		return err
	}
}
*/
