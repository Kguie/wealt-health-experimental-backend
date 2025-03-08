package employees

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmployee_Success(t *testing.T) {
	dto := CreateEmployeeDTO{
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: CustomDate{Time: time.Now().AddDate(-30, 0, 0)}, // 30 ans
		StartDate:   CustomDate{Time: time.Now()},
		Department:  "Sales",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
		},
	}

	err := ValidateEmployee(dto)
	assert.NoError(t, err)
}

func TestValidateEmployee_MissingFields(t *testing.T) {
	dto := CreateEmployeeDTO{}

	err := ValidateEmployee(dto)
	assert.Error(t, err)
}

func TestValidateEmployee_InvalidDateFormat(t *testing.T) {
	dto := CreateEmployeeDTO{
		FirstName:   "Jane",
		LastName:    "Smith",
		DateOfBirth: CustomDate{}, // Date vide
		StartDate:   CustomDate{},
		Department:  "HR",
		Address: Address{
			Street:  "456 Main St",
			City:    "San Francisco",
			State:   "CA",
			ZipCode: "94105",
		},
	}

	err := ValidateEmployee(dto)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "❌ DateOfBirth et StartDate doivent être au format 'dd-MM-yyyy' et non vides")
}

func TestValidateEmployee_InvalidDepartment(t *testing.T) {
	dto := CreateEmployeeDTO{
		FirstName:   "Alice",
		LastName:    "Brown",
		DateOfBirth: CustomDate{Time: time.Now().AddDate(-25, 0, 0)}, // 25 ans
		StartDate:   CustomDate{Time: time.Now()},
		Department:  "Finance", // Département non valide
		Address: Address{
			Street:  "789 Main St",
			City:    "Chicago",
			State:   "IL",
			ZipCode: "60601",
		},
	}

	err := ValidateEmployee(dto)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Department")
}

func TestValidateEmployee_InvalidZipCode(t *testing.T) {
	dto := CreateEmployeeDTO{
		FirstName:   "Bob",
		LastName:    "Williams",
		DateOfBirth: CustomDate{Time: time.Now().AddDate(-35, 0, 0)}, // 35 ans
		StartDate:   CustomDate{Time: time.Now()},
		Department:  "Engineering",
		Address: Address{
			Street:  "101 Broadway",
			City:    "Los Angeles",
			State:   "CA",
			ZipCode: "123", // Code postal invalide (doit être 5 chiffres)
		},
	}

	err := ValidateEmployee(dto)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ZipCode")
}
