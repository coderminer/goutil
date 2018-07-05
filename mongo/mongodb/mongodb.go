package db

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
)

const (
	dbhost    = "127.0.0.1:27017"
	authdb    = "admin"
	authuser  = "user"
	authpass  = "123456"
	timeout   = 60 * time.Second
	poollimit = 4096
)

var globalS *mgo.Session

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{dbhost},
		Timeout:   timeout,
		Source:    authdb,
		Username:  authuser,
		Password:  authpass,
		PoolLimit: poollimit,
	}

	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalf("Create Session: %s\n", err)
	}
	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	ms := globalS.Copy()
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

func IsEmpty(db, collection string) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func Count(db, collection string, query interface{}) (int, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Count()
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Insert(docs...)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).All(result)
}

func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func FindIter(db, collection string, query interface{}) *mgo.Iter {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Iter()
}

func Update(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Update(selector, update)
}

func Upsert(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.Upsert(selector, update)
	return err
}

func UpdateAll(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

func Remove(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Remove(selector)
}

func RemoveAll(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.RemoveAll(selector)
	return err
}

//insert one or multi documents
func BulkInsert(db, collection string, docs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Insert(docs...)
	return bulk.Run()
}

func BulkRemove(db, collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()

	bulk := c.Bulk()
	bulk.Remove(selector...)
	return bulk.Run()
}

func BulkRemoveAll(db, collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.RemoveAll(selector...)
	return bulk.Run()
}

func BulkUpdate(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Update(pairs...)
	return bulk.Run()
}

func BulkUpdateAll(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.UpdateAll(pairs...)
	return bulk.Run()
}

func BulkUpsert(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Upsert(pairs...)
	return bulk.Run()
}
