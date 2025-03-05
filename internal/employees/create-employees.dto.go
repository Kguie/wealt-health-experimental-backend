package employees

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-playground/validator"
)

// Format de date utilisé dans l'API
const dateFormat = "02-01-2006" // Format "dd-MM-yyyy"

// CustomDate permet de gérer la conversion JSON <-> date Go
type CustomDate struct {
	time.Time
}

// MarshalJSON : Convertit `time.Time` en `string` ("dd-MM-yyyy")
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(cd.Format(dateFormat))
}

// UnmarshalJSON : Convertit un `string` JSON en `time.Time`
func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	var dateStr string
	if err := json.Unmarshal(b, &dateStr); err != nil {
		return err
	}
	parsedTime, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return err
	}
	cd.Time = parsedTime
	return nil
}

// Address représente l'adresse d'un employé
type Address struct {
	Street  string `json:"street" validate:"required,min=2,max=100"`
	City    string `json:"city" validate:"required,min=2,max=50"`
	State   string `json:"state" validate:"required,oneof=AL AK AZ AR CA CO CT DE DC FL GA HI ID IL IN IA KS KY LA ME MD MA MI MN MS MO MT NE NV NH NJ NM NY NC ND OH OK OR PA RI SC SD TN TX UT VT VA WA WV WI WY"`
	ZipCode string `json:"zipCode" validate:"required,len=5,numeric"`
}

// CreateEmployeeDTO définit les données d'un employé avec validation
type CreateEmployeeDTO struct {
	FirstName   string     `json:"firstName" validate:"required,min=2,max=50"`
	LastName    string     `json:"lastName" validate:"required,min=2,max=50"`
	DateOfBirth CustomDate `json:"dateOfBirth" validate:"required"`
	StartDate   CustomDate `json:"startDate" validate:"required"`
	Department  string     `json:"department" validate:"required,oneof=Sales Marketing Engineering HR Legal"`
	Address     Address    `json:"address" validate:"required"`
}

// Création d'une instance unique du validateur
var validate = validator.New()

// ValidateEmployee valide les données et vérifie le format des dates
func ValidateEmployee(dto CreateEmployeeDTO) error {
	// Vérifier la structure avec go-playground/validator
	if err := validate.Struct(dto); err != nil {
		return err
	}

	// Vérifier si les dates sont bien définies (évite le "0001-01-01")
	if dto.DateOfBirth.IsZero() || dto.StartDate.IsZero() {
		return errors.New("❌ DateOfBirth et StartDate doivent être au format 'dd-MM-yyyy' et non vides")
	}

	return nil
}
