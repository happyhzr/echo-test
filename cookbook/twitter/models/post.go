package models

import (
	"github.com/insisthzr/echo-test/cookbook/twitter/conf"
	"github.com/insisthzr/echo-test/cookbook/twitter/db"

	"gopkg.in/mgo.v2/bson"
)

//Post post
type Post struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	From    string        `json:"to" bson:"to"`
	To      string        `json:"from" bson:"from"`
	Message string        `json:"message" bson:"message"`
}

// AddPost addpost
func (p *Post) AddPost() error {
	sess := db.NewDBSession()
	defer sess.Close()

	err := sess.DB("twitter").C("posts").Insert(p)
	if err != nil {
		return err
	}
	return nil
}

// FindPosts findposts
func FindPosts(to string, page int, limit int) ([]*Post, error) {
	posts := make([]*Post, 0)
	sess := db.NewDBSession()
	defer sess.Close()

	err := sess.DB(conf.DBName).C("posts").
		Find(bson.M{"to": to}).
		Skip(page * limit).
		Limit(limit).
		All(&posts)
	if err != nil {
		return posts, err
	}
	return posts, nil
}
