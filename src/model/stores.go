package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"mydb"
	"time"
)

type Store struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `bson:"name" json:"name"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func StoreCollectionName() string {
	return "local.store"
}

func (m Store) Insert() (err error) {
	mydb.Exec(StoreCollectionName(), func(c *mgo.Collection) {
		m.ID = bson.NewObjectId()
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
		err = c.Insert(m)
	})
	return
}

func FindStore(condition bson.M) (result *Store, err error) {
	mydb.Exec(StoreCollectionName(), func(c *mgo.Collection) {
		err = c.Find(condition).One(&result)
	})
	return
}

func FindStores(condition bson.M) (result []*Store, err error) {
	mydb.Exec(StoreCollectionName(), func(c *mgo.Collection) {
		err = c.Find(condition).All(&result)
	})
	return
}

func UpdateStore(selector, updater bson.M) (err error) {
	mydb.Exec(StoreCollectionName(), func(c *mgo.Collection) {
		err = c.Update(selector, updater)
	})
	return
}

func DeleteStore(condition bson.M) (err error) {
	mydb.Exec(StoreCollectionName(), func(c *mgo.Collection) {
		err = c.Remove(condition)
	})
	return
}
