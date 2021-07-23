package domain

import "github.com/EricOgie/egobank/reserrs"

// Client Personals
type Client struct {
	ID          string `json:"id" xml:"id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zip"`
	DateOfBirth string `json:"date_of_birth" xml:"DOB"`
	Status      string `json:"satus" xml:"status"`
}

// ClientRepo as the customer secondary Port
type ClientRepo interface {
	FindAllClient() ([]*Client, *reserrs.MyError)
	ClientByID(string) (*Client, *reserrs.MyError)
}
