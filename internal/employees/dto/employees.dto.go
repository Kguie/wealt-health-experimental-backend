package employees

import (
	"time"

	"github.com/go-playground/validator"
)

// Address représente l'adresse d'un employé
type Address struct {
	Street  string `json:"street" validate:"required,min=2,max=100"`
	City    string `json:"city" validate:"required,min=2,max=50"`
	State   string `json:"state" validate:"required,oneof=AL AK AZ AR CA CO CT DE DC FL GA HI ID IL IN IA KS KY LA ME MD MA MI MN MS MO MT NE NV NH NJ NM NY NC ND OH OK OR PA RI SC SD TN TX UT VT VA WA WV WI WY"`
	ZipCode string `json:"zipCode" validate:"required,len=5,numeric"`
}

// EmployeeDTO définit les données d'un employé avec validation
type EmployeeDTO struct {
	FirstName   string    `json:"firstName" validate:"required,min=2,max=50"`
	LastName    string    `json:"lastName" validate:"required,min=2,max=50"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
	StartDate   time.Time `json:"startDate" validate:"required"`
	Department  string    `json:"department" validate:"required,oneof=Sales Marketing Engineering HR Legal"`
	Address     Address   `json:"address" validate:"required"` // Struct imbriquée validée
}

// Création d'une instance unique du validateur
var validate = validator.New()

// Fonction pour valider les données d'un DTO
func ValidateEmployee(dto EmployeeDTO) error {
	return validate.Struct(dto)
}
