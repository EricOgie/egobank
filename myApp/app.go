package myapp

import (
	"log"
	"net/http"

	"github.com/EricOgie/egobank/domain"
	"github.com/EricOgie/egobank/konstants"
	"github.com/EricOgie/egobank/service"
	"github.com/gorilla/mux"
)

//StartApp function to init App
func StartApp() {

	router := mux.NewRouter()

	//Wiring Connection
	// handler := ClientHandler{service.CreateNewClientServiceCore(domain.CreateNewClientStub())} for stub test
	handler := ClientHandler{service.CreateNewClientServiceCore(domain.CreateNewRepoDBMysql())}

	router.HandleFunc("/", greet).Methods(http.MethodGet)
	router.HandleFunc("/clients", handler.indexClient).Methods(http.MethodGet)
	router.HandleFunc("/clients/{id:[0-9]+}", handler.findOneClient).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(konstants.PORT, router))
}
