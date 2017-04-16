package ServerHandler

import (
	"ShuffleEat/Model/DataBase"
	"encoding/json"
	"io"
	"log"
)

const dbName = "Recepies"

// NewRecepieHandler handle new recepie
type NewRecepieHandler struct {
	dbConnector *DataBase.Connector
}

// FactoryNewRecepieHandler creates new instance
func FactoryNewRecepieHandler() *NewRecepieHandler {
	this := new(NewRecepieHandler)
	dbConnector := new(DataBase.Connector)
	dbConnector.InitDatabase(dbName)
	this.dbConnector = dbConnector
	return this
}

type test_struct struct {
	Test string
}

// ProceedData unmarshal json and add new recepie
func (handler *NewRecepieHandler) ProceedData(data io.ReadCloser) {
	decoder := json.NewDecoder(data)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer data.Close()
	log.Println(t.Test)
}
