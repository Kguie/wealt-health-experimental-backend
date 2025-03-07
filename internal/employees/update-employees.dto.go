package employees

import (
	"errors"
)

type address struct {
	Street  *string `json:"street,omitempty" validate:"omitempty,min=2,max=100"`
	City    *string `json:"city,omitempty" validate:"omitempty,min=2,max=50"`
	State   *string `json:"state,omitempty" validate:"omitempty,oneof=AL AK AZ AR CA CO CT DE DC FL GA HI ID IL IN IA KS KY LA ME MD MA MI MN MS MO MT NE NV NH NJ NM NY NC ND OH OK OR PA RI SC SD TN TX UT VT VA WA WV WI WY"`
	ZipCode *string `json:"zipCode,omitempty" validate:"omitempty,len=5,numeric"`
}

// CreateEmployeeDTO définit les données d'un employé avec validation
type UpdateEmployeeDTO struct {
	DateOfBirth *CustomDate `json:"dateOfBirth,omitempty"`
	StartDate   *CustomDate `json:"startDate,omitempty"`
	Department  *string     `json:"department,omitempty" validate:"omitempty,oneof=Sales Marketing Engineering HR Legal"`
	Address     *address    `json:"address,omitempty"`
}

func ValidateEmployeeUpdate(dto UpdateEmployeeDTO) error {
	// Vérifier la structure avec go-playground/validator
	if err := validate.Struct(dto); err != nil {
		return err
	}

	// Vérifier si les dates sont bien définies (évite le "0001-01-01")
	if (dto.DateOfBirth != nil && dto.DateOfBirth.IsZero()) || (dto.StartDate != nil && dto.StartDate.IsZero()) {
		return errors.New("❌ DateOfBirth et StartDate doivent être au format 'dd-MM-yyyy'")
	}

	return nil
}
