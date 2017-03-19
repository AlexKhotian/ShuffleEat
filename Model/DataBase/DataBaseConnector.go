package DataBase

import (
	"log"

	"gopkg.in/mgo.v2"
)

// IConnector interface to database
type IConnector interface {
	InitDatabase()
	AddItems(collection string, items ...interface{})
	RemoveItems(collection string, selector interface{})
	ModifyItems(collection string, selector interface{}, items ...interface{})
	FindItems(collection string, query interface{}) []interface{}
	Shutdown()
}

// Connector connects to database
type Connector struct {
	dbSession *mgo.Session
	database  *string
}

// InitDatabase init db with name
func (conn *Connector) InitDatabase(database string) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	conn.dbSession = session.Copy()
	conn.dbSession.SetMode(mgo.Monotonic, true)
	conn.database = &database
}

// Shutdown close connection to db
func (conn *Connector) Shutdown() {
	conn.dbSession.Close()
}

// AddItems to selected collection
func (conn *Connector) AddItems(collection string, items ...interface{}) {
	c := conn.dbSession.DB(*conn.database).C(collection)
	err := c.Insert(&items)
	if err != nil {
		log.Println(err)
	}
}

// RemoveItems from selected collection
func (conn *Connector) RemoveItems(collection string, selector interface{}) {
	c := conn.dbSession.DB(*conn.database).C(collection)
	err := c.Remove(&selector)
	if err != nil {
		log.Println(err)
	}
}

// ModifyItems in selected collection
func (conn *Connector) ModifyItems(collection string, selector interface{}, items ...interface{}) {
	c := conn.dbSession.DB(*conn.database).C(collection)
	for _, item := range items {
		err := c.Update(selector, item)
		if err != nil {
			log.Println(err, item)
		}
	}
}

// FindItems in selected collection with provided query
func (conn *Connector) FindItems(collection string, query interface{}) []interface{} {
	c := conn.dbSession.DB(*conn.database).C(collection)
	var result []interface{}
	err := c.Find(query).All(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}
