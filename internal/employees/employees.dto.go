package employees

import "time"

type EmployeeDTO struct {
	Id          string          `json:"id"`
	FirstName   string          `json:"firstName"`
	LastName    string          `json:"lastName"`
	DateOfBirth time.Time       `json:"dateOfBirth"`
	StartDate   time.Time       `json:"startDate"`
	Department  string          `json:"department"`
	Address     EmployeeAddress `json:"address"`
}

type EmployeeAddress struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}
