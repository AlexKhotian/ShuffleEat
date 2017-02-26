package ServerHandler

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// HTTPHandlerUtil handles http requsets
type HTTPHandlerUtil interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// HTTPHandlerUtilImpl implement interface
type HTTPHandlerUtilImpl struct {
}

// HTTPHandlerFactory creates a interface for HTTPHandlerUtil
func HTTPHandlerFactory() HTTPHandlerUtil {
	thisHandler := new(HTTPHandlerUtilImpl)
	return thisHandler
}

func (handler *HTTPHandlerUtilImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "" {
		path = "/mainView.html"
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

func (handler *HTTPHandlerUtilImpl) parseRequest(path *string) (string, string) {
	var contentType string
	modifiedSourcePath := "../../Webview" + *(path)
	if strings.HasSuffix(*path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(*path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(*path, ".js") {
		contentType = "application/javascript"
		modifiedSourcePath = "ViewScripts" + *(path)
	} else if strings.HasSuffix(*path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}
	return modifiedSourcePath, contentType
}
