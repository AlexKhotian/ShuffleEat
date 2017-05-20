package ServerHandler

import "ShuffleEat/Model/DataBase"

// RandomMealHandler handles request for random meal,
// Gets random recipe from database
// Converts it to RecipeFrom responce
// Write to responce
type RandomMealHandler struct {
	dbConnector *DataBase.Connector
}

// FactoryRandomMealHandler creates new handler
func FactoryRandomMealHandler() *RandomMealHandler {
	this := new(RandomMealHandler)
	dbConnector := new(DataBase.Connector)
	//dbConnector.InitDatabase(dbName)
	this.dbConnector = dbConnector
	return this
}

//
