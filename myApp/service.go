package myapp

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/EricOgie/egobank/domain"
	"github.com/EricOgie/egobank/konstants"
)

func serverClients(clientList []*domain.Client, res http.ResponseWriter, req *http.Request) {
	if req.Header.Get(konstants.ContentType) == konstants.TypeXML {
		res.Header().Add(konstants.ContentType, konstants.TypeXML)
		xml.NewEncoder(res).Encode(clientList)
	} else {
		res.Header().Add(konstants.ContentType, konstants.TypeJSON)
		json.NewEncoder(res).Encode(clientList)

	}
}
