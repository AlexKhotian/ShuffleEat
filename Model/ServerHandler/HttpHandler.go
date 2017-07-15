package ServerHandler

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// HTTPHandlerUtil implement interface
type HTTPHandlerUtil struct {
	newRecepieHandler      *NewRecepieHandler
	randomMealHandler      *RandomMealHandler
	searchCategorieHandler *CategorieSearchHandler
}

// HTTPHandlerFactory creates a interface for HTTPHandlerUtil
func HTTPHandlerFactory() *HTTPHandlerUtil {
	thisHandler := new(HTTPHandlerUtil)
	thisHandler.newRecepieHandler = FactoryNewRecepieHandler()
	thisHandler.randomMealHandler = FactoryRandomMealHandler()
	thisHandler.searchCategorieHandler = FactoryCategorieSearchHandler()
	return thisHandler
}

func (handler *HTTPHandlerUtil) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			if r.URL.Path == "/NewRecipe" {
				handler.newRecepieHandler.ProceedData(r.Body)
				return
			} else if r.URL.Path == "/PickCategorieRequest" {
				handler.searchCategorieHandler.SearchForRecipes(r.Body, w)
				return
			}
		}
	}

	if r.Method == "GET" {
		if r.URL.Path == "/RandomMeal" {
			handler.randomMealHandler.GenerateRandomMeal(w)
			return
		}
	}
	path := r.URL.Path
	if path == "/" {
		path = "/StartPage.html"
	}

	modifiedPath, contentType := handler.parseRequest(&path)
	data, err := ioutil.ReadFile(string(modifiedPath))
	log.Println(modifiedPath)
	if err == nil {
		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 not found"))
	}
}

func (handler *HTTPHandlerUtil) parseRequest(path *string) (string, string) {
	var contentType string
	modifiedSourcePath := "../../Webview" + *(path)
	if strings.HasSuffix(*path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(*path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(*path, ".js") {
		modifiedSourcePath = "../../ViewModel" + *(path)
		contentType = "application/javascript"
	} else if strings.HasSuffix(*path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}
	return modifiedSourcePath, contentType
}
