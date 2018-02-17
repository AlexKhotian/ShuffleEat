package ServerHandler

import (
	"ShuffleEat/Model/DataBase"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// CategorieSearchHandler handle search recipe
type CategorieSearchHandler struct {
	dbConnector *DataBase.Connector
}

// FactoryCategorieSearchHandler creates new instance
func FactoryCategorieSearchHandler() *CategorieSearchHandler {
	this := new(CategorieSearchHandler)
	dbConnector := new(DataBase.Connector)
	if err := dbConnector.InitDatabase(dbName); err != nil {
		return nil
	}
	this.dbConnector = dbConnector
	return this
}

// SearchForRecipes converts json, search in DB, returns result to ResponseWriter
func (handler *CategorieSearchHandler) SearchForRecipes(data io.ReadCloser, w http.ResponseWriter) {
	decoder := json.NewDecoder(data)
	var searchRequest CategorieSearchRequest
	var dataR []byte
	data.Read(dataR)
	err := decoder.Decode(&searchRequest)
	log.Println(dataR)
	if err != nil {
		log.Println("FAILED to parse request with err:", err)
	}
	data.Close()
	interfaceVars := handler.dbConnector.FindItems(
		dbCollection, bson.M{"_categorieType": DataBase.Categorie(searchRequest.CategorieType)})

	recipeTitles := &CategorieSearchResponse{}
	for _, value := range interfaceVars {
		log.Println(value.RecipeTitle)
		recipeTitles.RecipeTitles = append(recipeTitles.RecipeTitles, value.RecipeTitle)
	}
	if len(recipeTitles.RecipeTitles) != 0 {
		response, errM := json.Marshal(&recipeTitles)
		if errM != nil {
			log.Println(errM)
			return
		}
		w.Write(response)
	} else {
		log.Println("SearchForRecipes: no values found")
	}
}
