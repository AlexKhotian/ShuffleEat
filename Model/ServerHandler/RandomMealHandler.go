package ServerHandler

import (
	"ShuffleEat/Model/DataBase"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

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
	dbConnector.InitDatabase(dbName)
	this.dbConnector = dbConnector
	return this
}

// GenerateRandomMeal finds a random meal and returns as JSON http response
func (handler *RandomMealHandler) GenerateRandomMeal(w http.ResponseWriter) {
	pseudoRandomID := rand.Intn(int(handler.dbConnector.GetCollectionSize(dbCollection)))
	log.Println("R_Id:", pseudoRandomID)
	interfaceVars := handler.dbConnector.FindItems(dbCollection, bson.M{"id": pseudoRandomID})
	var response []byte
	for _, value := range interfaceVars {
		log.Println("Found empty interface")
		result, isRightType := value.(DataBase.Recipe)
		if isRightType {
			responseRecipe := &RecipeForm{}
			log.Println("Found:", responseRecipe.RecipeTitle)
			responseRecipe.Description = result.Description
			responseRecipe.Categorie = uint16(result.CategorieType)
			responseRecipe.Ingredients = result.Ingredients
			responseRecipe.Quantities = result.Quantities
			responseRecipe.RecipeTitle = result.RecipeTitle
			jsonBytes, err := json.Marshal(responseRecipe)
			if err != nil {
				log.Println(err)
				continue
			}
			response = append(response, jsonBytes...)
		}
	}
	if len(response) != 0 {
		w.Write(response)
	} else {
		log.Println("GenerateRandomMeal: no new values")
	}
}
