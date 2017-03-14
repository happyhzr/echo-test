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
	Password  string        `json:"password" bson:"password"`
	Token     string        `json:"token" bson:"-"`
	Followers []string      `json:"followers" bson:"followers"`
}

// AddUser adduser
func (u *User) AddUser() error {
	sess := db.NewDBSession()
	defer sess.Close()

	selector := bson.M{"email": u.Email}
	_, err := sess.DB(conf.DBName).C("users").Upsert(selector, u)
	return err
}

// ValidUser validuser
func (u *User) ValidUser() (bool, error) {
	sess := db.NewDBSession()
	defer sess.Close()

	queryUser := new(User)
	selector := bson.M{"email": u.Email}
	err := sess.DB(conf.DBName).C("users").Find(selector).One(queryUser)
	if err == mgo.ErrNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	if u.Password == queryUser.Password {
		return true, nil
	}
	return false, nil
}

// UserExist check exist
func UserExist(email string) (bool, error) {
	sess := db.NewDBSession()
	defer sess.Close()

	selector := bson.M{"email": email}
	count, err := sess.DB(conf.DBName).C("users").Find(selector).Count()
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
	err := sess.DB(conf.DBName).C("users").UpdateId(id, update)
	if err != nil {
		return err
	}
	return nil
}
