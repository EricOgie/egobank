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
	clients, err := cHandler.service.GetAllClient()
	if err != nil {
		res.WriteHeader(err.Code)
		fmt.Fprintf(res, err.Message)
	} else {
		serverClients(clients, res, req)
	}

}

func (cHandler ClientHandler) findOneClient(res http.ResponseWriter, req *http.Request) {
	var clientList = make([]*domain.Client, 0)
	reqVars := mux.Vars(req)
	id := reqVars["id"]
	client, err := cHandler.service.GetOneClient(id)
	if err != nil {
		res.WriteHeader(err.Code)
		fmt.Fprintf(res, err.Message)
	} else {
		clientList = append(clientList, client)
		serverClients(clientList, res, req)
	}

}
