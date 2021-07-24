package service

import (
	"github.com/EricOgie/egobank/domain"
	"github.com/EricOgie/egobank/reserrs"
)

// ClientService as a primary port
type ClientService interface {
	GetAllClient(string) ([]*domain.Client, *reserrs.MyError)
	GetOneClient(string) (*domain.Client, *reserrs.MyError)
}

//ClientServiceCore as the application core for Client biz logic
//with clientRepo injection
type ClientServiceCore struct {
	clientRepo domain.ClientRepo
}

// GetAllClient method of ClienService Port Inplementation with ClientServiceCore
func (core ClientServiceCore) GetAllClient(status string) ([]*domain.Client, *reserrs.MyError) {
	return core.clientRepo.FindAllClient(status)
}

// GetOneClient method of ClienService Port Inplementation with ClientServiceCore
func (core ClientServiceCore) GetOneClient(id string) (*domain.Client, *reserrs.MyError) {
	return core.clientRepo.ClientByID(id)
}

// CreateNewClientServiceCore as helper function to instantiate clientServiceCore
// CreateNewClientServiceCore will take the dependecy of CLientRepo so it can
// be injected into the ClientSerivceCore
func CreateNewClientServiceCore(repo domain.ClientRepo) ClientServiceCore {
	return ClientServiceCore{repo}
}
