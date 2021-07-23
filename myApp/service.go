package myapp

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/EricOgie/egobank/konstants"
)

func serveResponse(response interface{}, status int, res http.ResponseWriter, req *http.Request) {
	if req.Header.Get(konstants.ContentType) == konstants.TypeXML {
		res.Header().Add(konstants.ContentType, konstants.TypeXML)
		res.WriteHeader(status)
		if err := xml.NewEncoder(res).Encode(response); err != nil {
			panic(err)
		}

	} else {
		res.Header().Add(konstants.ContentType, konstants.TypeJSON)
		res.WriteHeader(status)
		if err := json.NewEncoder(res).Encode(response); err != nil {
			panic(err)
		}

	}

}
