package ServerHandler

import "net/http"

// IServerRoutine starts a web handler server
type IServerRoutine interface {
	RunServer()
}

// ServerRoutineImpl implementation for ServerRoutine interface
// contains handler
type ServerRoutineImpl struct {
	httpHandlerUtil *HTTPHandlerUtil
}

// ServerRoutineFactory creats intance of ServerRoutine
func ServerRoutineFactory() (IServerRoutine, error) {
	thisServer := new(ServerRoutineImpl)
	var err error
	thisServer.httpHandlerUtil, err = HTTPHandlerFactory()
	if err != nil {
		return nil, err
	}
	return thisServer, nil
}

// RunServer starts a web server
func (serverImpl *ServerRoutineImpl) RunServer() {
	http.Handle("/", serverImpl.httpHandlerUtil)
	http.ListenAndServe(":8080", nil)
}
