package service

import (
	"github.com/EricOgie/egobank/domain"
	"github.com/EricOgie/egobank/dto"
	"github.com/EricOgie/egobank/reserrs"
)

// ClientService as a primary port
type ClientService interface {
	GetAllClient(string) ([]*dto.ClientResponse, *reserrs.MyError)
	GetOneClient(string) (*dto.ClientResponse, *reserrs.MyError)
}

//ClientServiceCore as the application core for Client biz logic
//with clientRepo injection
type ClientServiceCore struct {
	clientRepo domain.ClientRepo
}

// GetAllClient method of ClienService Port Inplementation with ClientServiceCore
func (core ClientServiceCore) GetAllClient(status string) ([]*dto.ClientResponse, *reserrs.MyError) {
	clientList, err := core.clientRepo.FindAllClient(status)
	var responseDtos = make([]*dto.ClientResponse, 0)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(clientList); i++ {
		client := clientList[i]
		clientDto := client.ToDto()
		responseDtos = append(responseDtos, &clientDto)
	}

	return responseDtos, nil

}

// GetOneClient method of ClienService Port Inplementation with ClientServiceCore
func (core ClientServiceCore) GetOneClient(id string) (*dto.ClientResponse, *reserrs.MyError) {
	client, err := core.clientRepo.ClientByID(id)
	if err != nil {
		return nil, err
	}
	response := client.ToDto()
	return &response, nil

}

// CreateNewClientServiceCore as helper function to instantiate clientServiceCore
// CreateNewClientServiceCore will take the dependecy of CLientRepo so it can
// be injected into the ClientSerivceCore
func CreateNewClientServiceCore(repo domain.ClientRepo) ClientServiceCore {
	return ClientServiceCore{repo}
}
