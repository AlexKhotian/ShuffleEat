package ServerHandler

import (
	"ShuffleEat/Model/DataBase"
	"encoding/json"
	"io"
)

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
	var recipeForm RecipeForm
	err := decoder.Decode(&recipeForm)
	if err != nil {
		panic(err)
	}
	data.Close()

	databaseRecipe := convertPostRequestToDatabaseFormat(&recipeForm)
	handler.dbConnector.AddItems(dbCollection, databaseRecipe)
}

func convertPostRequestToDatabaseFormat(requestData *RecipeForm) *DataBase.Recipe {
	var recipe DataBase.Recipe
	recipe.Description = requestData.Description
	recipe.Ingredients = requestData.Ingredients
	recipe.Quantities = requestData.Quantities
	recipe.CategorieType = DataBase.Categorie(requestData.Categorie)
	recipe.RecipeTitle = requestData.RecipeTitle
	return &recipe
}
