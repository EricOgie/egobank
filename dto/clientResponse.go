package dto

// ClientResponse Personals
type ClientResponse struct {
	ID          string `json:"customer_id" xml:"id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zip"`
	DateOfBirth string `json:"date_of_birth" xml:"DOB"`
	Status      string `json:"satus" xml:"status"`
}
