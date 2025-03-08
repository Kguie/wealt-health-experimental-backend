package employees

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmployeeUpdate_Success(t *testing.T) {
	newDept := "Marketing"
	newCity := "Boston"
	dto := UpdateEmployeeDTO{
		DateOfBirth: &CustomDate{Time: time.Now().AddDate(-25, 0, 0)}, // 25 ans
		StartDate:   &CustomDate{Time: time.Now()},
		Department:  &newDept,
		Address: &address{
			City: &newCity,
		},
	}

	err := ValidateEmployeeUpdate(dto)
	assert.NoError(t, err)
}

func TestValidateEmployeeUpdate_EmptyFields(t *testing.T) {
	dto := UpdateEmployeeDTO{}

	err := ValidateEmployeeUpdate(dto)
	assert.NoError(t, err) // Comme tout est `omitempty`, une requête vide ne doit pas renvoyer d'erreur
}

func TestValidateEmployeeUpdate_InvalidDate(t *testing.T) {
	newDept := "HR"
	dto := UpdateEmployeeDTO{
		DateOfBirth: &CustomDate{}, // Date vide (0001-01-01)
		StartDate:   &CustomDate{},
		Department:  &newDept,
	}

	err := ValidateEmployeeUpdate(dto)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "❌ DateOfBirth et StartDate doivent être au format 'dd-MM-yyyy'")
}

func TestValidateEmployeeUpdate_InvalidDepartment(t *testing.T) {
	invalidDept := "Finance" // Non inclus dans la liste des départements valides
	dto := UpdateEmployeeDTO{
		Department: &invalidDept,
	}

	err := ValidateEmployeeUpdate(dto)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Department")
}

func TestValidateEmployeeUpdate_InvalidZipCode(t *testing.T) {
	newZip := "123" // Doit être exactement 5 chiffres
	dto := UpdateEmployeeDTO{
		Address: &address{
			ZipCode: &newZip,
		},
	}

	err := ValidateEmployeeUpdate(dto)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ZipCode")
}
