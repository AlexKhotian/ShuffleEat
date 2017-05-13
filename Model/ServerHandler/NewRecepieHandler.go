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
	Quantities  string `json:"_quantities"`
	Description string `json:"_description"`
	Categorie   uint16 `json:"_categorie"`
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
	//dbConnector.InitDatabase(dbName)
	this.dbConnector = dbConnector
	return this
}

// ProceedData unmarshal json and add new recepie
func (handler *NewRecepieHandler) ProceedData(data io.ReadCloser) {
	decoder := json.NewDecoder(data)
	var recipeForm RecipeForm
	err := decoder.Decode(&recipeForm)
	if err != nil {
		panic(err)
	}
	data.Close()
	log.Println(recipeForm)
}

func convertPostRequestToDatabaseFormat(requestData *RecipeForm) *DataBase.Recipe {
	var recipe DataBase.Recipe
	//	recipe.

	return &recipe
}

func parseIngredients(ingredientString *string) *[]DataBase.Ingredient {
	return nil
}
