package ServerHandler

import (
	"ShuffleEat/Model/DataBase"
	"encoding/json"
	"io"
	"log"
)

// RecipeForm contains in about recipe from html form
type RecipeForm struct {
	RecipeTitle string `json:"_title"`
	Ingredients string `json:"_ingredients"`
	Description string `json:"_description"`
	Categorie   string `json:"_categorie"`
}

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

// ProceedData unmarshal json and add new recepie
func (handler *NewRecepieHandler) ProceedData(data io.ReadCloser) {
	decoder := json.NewDecoder(data)
	var recipe RecipeForm
	err := decoder.Decode(&recipe)
	if err != nil {
		panic(err)
	}
	defer data.Close()
	log.Println(recipe)
}
