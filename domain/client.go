package domain

import (
	"github.com/EricOgie/egobank/dto"
	"github.com/EricOgie/egobank/konstants"
	"github.com/EricOgie/egobank/reserrs"
)

// Client Personals
type Client struct {
	ID          string `db:"customer_id" xml:"id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zip"`
	DateOfBirth string `db:"date_of_birth" xml:"DOB"`
	Status      string `json:"satus" xml:"status"`
}

// ToDto func to parse Client to DTO object
func (client Client) ToDto() dto.ClientResponse {
	return dto.ClientResponse{
		ID:          client.ID,
		Name:        client.Name,
		City:        client.City,
		Zipcode:     client.Zipcode,
		DateOfBirth: client.DateOfBirth,
		Status:      konstants.GetClientStatus(client.Status),
	}

}

// ClientRepo as the customer secondary Port
type ClientRepo interface {
	FindAllClient(string) ([]*Client, *reserrs.MyError)
	ClientByID(string) (*Client, *reserrs.MyError)
}
