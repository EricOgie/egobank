package myapp

import (
	"fmt"
	"net/http"

	"github.com/EricOgie/egobank/domain"
	"github.com/EricOgie/egobank/service"
	"github.com/gorilla/mux"
)

// ClientHandler that will implement the ClientService
type ClientHandler struct {
	service service.ClientService
}

func greet(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Welcome To Hello World, GO</h1>")
}

func (cHandler ClientHandler) indexClient(res http.ResponseWriter, req *http.Request) {
	status := req.URL.Query().Get("status")
	clients, err := cHandler.service.GetAllClient(status)
	if err != nil {
		serveResponse(err.ErrorMsg(), err.Code, res, req)
	} else {
		serveResponse(clients, http.StatusOK, res, req)
	}

}

func (cHandler ClientHandler) findOneClient(res http.ResponseWriter, req *http.Request) {
	var clientList = make([]*domain.Client, 0)
	reqVars := mux.Vars(req)
	id := reqVars["id"]
	client, err := cHandler.service.GetOneClient(id)
	if err != nil {
		serveResponse(err.ErrorMsg(), err.Code, res, req)
	} else {
		clientList = append(clientList, client)
		serveResponse(clientList, http.StatusOK, res, req)
	}

}
