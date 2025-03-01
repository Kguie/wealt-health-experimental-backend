package employees

import (
	"fmt"
	"net/http"
	"wealth-health-backend/ent"
)

type EmployeeHandler struct {
	Client *ent.Client
}

// NewEmployeeHandler crée une instance de `EmployeeHandler`
func NewEmployeeHandler(client *ent.Client) *EmployeeHandler {
	return &EmployeeHandler{Client: client}
}

func (employee *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an Employee")
}

func (employee *EmployeeHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all Employees")
}

func (employee *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an Employee by ID")
}

func (employee *EmployeeHandler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an Employee by ID")
}

func (employee *EmployeeHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an Employee by ID")
}
