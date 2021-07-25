package myapp

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EricOgie/egobank/domain"
	"github.com/EricOgie/egobank/konstants"
	"github.com/EricOgie/egobank/mylogger"
	"github.com/EricOgie/egobank/service"
	"github.com/gorilla/mux"
)

//StartApp function to init App
func StartApp() {

	setAppEnvs()
	router := mux.NewRouter()

	//Wiring Connection
	// handler := ClientHandler{service.CreateNewClientServiceCore(domain.CreateNewClientStub())} for stub test
	handler := ClientHandler{service.CreateNewClientServiceCore(domain.CreateNewRepoDBMysql())}

	router.HandleFunc("/", greet).Methods(http.MethodGet)
	router.HandleFunc("/clients", handler.indexClient).Methods(http.MethodGet)
	router.HandleFunc("/clients/{id:[0-9]+}", handler.findOneClient).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func setAppEnvs() {
	os.Setenv(konstants.ServerKey, "localhost")
	os.Setenv(konstants.PortKey, "8000")
	os.Setenv(konstants.DBUserKey, konstants.DBUser)
	os.Setenv(konstants.DBAddKey, konstants.DBAddress)
	os.Setenv(konstants.DBPortKEy, konstants.DBPort)
	os.Setenv(konstants.DBNameKey, konstants.DBName)
	mylogger.Info("Env Set ...")

	if !sanityCheck() {
		mylogger.Error("Envronment Variable not set")
		return
	}

}

func sanityCheck() bool {
	if os.Getenv("SERVER_PORT") == "" || os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("DB_USER") == "" || os.Getenv("DB_ADDR") == "" ||
		os.Getenv("DB_PORT") == "" || os.Getenv("DB_NAME") == "" {

		return false
	}
	return true
}
