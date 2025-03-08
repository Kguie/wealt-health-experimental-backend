package employees

import (
	"context"
	"testing"
	"time"
	"wealth-health-backend/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func parseToDate(stringDate string) time.Time {
	layout := "02-01-2006"
	parsedTime, err := time.Parse(layout, stringDate)
	if err != nil {
		return time.Now()
	}
	return parsedTime
}

var testCreateEmployeeDTO = CreateEmployeeDTO{
	FirstName:   "John",
	LastName:    "Doe",
	DateOfBirth: CustomDate{parseToDate("01-01-1895")},
	StartDate:   CustomDate{parseToDate("01-01-2001")},
	Department:  "Legal",
	Address: Address{
		Street:  "123 Main St",
		City:    "Paris",
		State:   "FR",
		ZipCode: "75000",
	},
}

func TestCreateEmployee(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()

	handler := NewEmployeeHandler(client)
	resp, err := handler.Create(ctx, testCreateEmployeeDTO)
	assert.NoError(t, err)
	assert.Equal(t, "L'employé a été ajouté avec succès.", resp["message"])

}

func TestListEmployees(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	handler := NewEmployeeHandler(client)
	ctx := context.Background()

	// Ajouter un employé pour tester la récupération
	_, _ = handler.Create(ctx, testCreateEmployeeDTO)

	employees, err := handler.List(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, employees)
	assert.Equal(t, "John", (*employees)[0].FirstName)
}

func TestGetEmployeeByID(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	handler := NewEmployeeHandler(client)
	ctx := context.Background()

	// Ajouter un employé
	emp, _ := client.Employee.Create().
		SetFirstName("Jane").
		SetLastName("Doe").
		SetDateOfBirth(time.Now().AddDate(-25, 0, 0)).
		SetStartDate(time.Now()).
		SetDepartment("HR").
		SetStreet("456 Main St").
		SetCity("Lyon").
		SetState("FR").
		SetZipCode("69000").
		Save(ctx)

	// Récupérer l'employé par ID
	res, err := handler.GetByID(ctx, emp.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Jane", res.FirstName)
	assert.Equal(t, "HR", res.Department)
}

func TestUpdateEmployeeByID(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	handler := NewEmployeeHandler(client)
	ctx := context.Background()

	// Ajouter un employé
	emp, _ := client.Employee.Create().
		SetFirstName("Paul").
		SetLastName("Smith").
		SetDateOfBirth(time.Now().AddDate(-40, 0, 0)).
		SetStartDate(time.Now()).
		SetDepartment("Legal").
		SetStreet("789 Main St").
		SetCity("Marseille").
		SetState("NY").
		SetZipCode("13000").
		Save(ctx)

	updateDTO := UpdateEmployeeDTO{
		Department: stringPointer("Engineering"),
	}

	updatedEmp, err := handler.UpdateByID(ctx, emp.ID, updateDTO)
	assert.NoError(t, err)
	assert.Equal(t, "Engineering", updatedEmp.Department)

}

func TestDeleteEmployeeByID(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	handler := NewEmployeeHandler(client)
	ctx := context.Background()

	// Ajouter un employé
	emp, _ := client.Employee.Create().
		SetFirstName("Alice").
		SetLastName("Brown").
		SetDateOfBirth(time.Now().AddDate(-35, 0, 0)).
		SetStartDate(time.Now()).
		SetDepartment("Support").
		SetStreet("101 Main St").
		SetCity("Bordeaux").
		SetState("FR").
		SetZipCode("33000").
		Save(ctx)

	// Supprimer l'employé
	resp, err := handler.DeleteByID(ctx, emp.ID)
	assert.NoError(t, err)
	assert.Equal(t, "L'employé a été supprimé avec succès.", resp["message"])

	// Vérifier que l'employé n'existe plus
	_, err = handler.GetByID(ctx, emp.ID)
	assert.Error(t, err)
}

// Fonction utilitaire pour gérer les pointeurs de string
func stringPointer(s string) *string {
	return &s
}
