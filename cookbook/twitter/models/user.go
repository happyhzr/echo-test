package models

import (
	"github.com/insisthzr/echo-test/cookbook/twitter/conf"
	"github.com/insisthzr/echo-test/cookbook/twitter/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User user
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password,omitempty" bson:"password"`
	Token     string        `json:"token,omitempty" bson:"-"`
	Followers []string      `json:"followers,omitempty" bson:"followers"`
}

// AddUser adduser
func (u *User) AddUser() error {
	sess := db.NewDBSession()
	defer sess.Close()

	err := sess.DB(conf.DB_NAME).C("users").Insert(u)
	return err
}

// FindUserByEmail FindUserByEmail
func FindUserByEmail(email string) (*User, error) {
	sess := db.NewDBSession()
	defer sess.Close()

	existUser := new(User)
	selector := bson.M{"email": email}
	err := sess.DB(conf.DB_NAME).C("users").Find(selector).One(existUser)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return existUser, nil
}

// UserExist check exist
func UserExist(email string) (bool, error) {
	sess := db.NewDBSession()
	defer sess.Close()

	selector := bson.M{"email": email}
	count, err := sess.DB(conf.DB_NAME).C("users").Find(selector).Count()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// AddFollower add follower
func AddFollower(id bson.ObjectId, followerID string) error {
	sess := db.NewDBSession()
	defer sess.Close()

	update := bson.M{"$addToSet": bson.M{"followers": followerID}}
	err := sess.DB(conf.DB_NAME).C("users").UpdateId(id, update)
	if err != nil {
		return err
	}
	return nil
}
