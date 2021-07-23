package domain

// ClientRepoStub that will implement the ClientRepo interface
type ClientRepoStub struct {
	Clients []Client
}

// FindAllClient interface function implemtation on ClientStub
func (cRStub ClientRepoStub) FindAllClient() ([]Client, error) {
	return cRStub.Clients, nil
}

// CreateNewClientStub helper function to instantiate new ClientStub
func CreateNewClientStub() ClientRepoStub {
	client := []Client{
		{"100", "Osamuyi", "Benin", "234", "18-02-88", "1"},
		{"101", "Justin", "Benin", "234", "18-02-92", "0"},
		{"102", "Dominic", "Lagos", "234", "18-02-96", "1"},
	}
	return ClientRepoStub{client}
}
