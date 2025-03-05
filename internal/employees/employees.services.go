package employees

import (
	"encoding/json"
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
	ctx := r.Context()

	// 🔥 Récupérer tous les employés depuis la BDD
	employees, err := employee.Client.Employee.Query().All(ctx)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des employés", http.StatusInternalServerError)
		return
	}

	var employeesDTO []EmployeeDTO
	for _, emp := range employees {
		employeesDTO = append(employeesDTO, EmployeeDTO{
			Id:          emp.ID.String(),
			FirstName:   emp.FirstName,
			LastName:    emp.LastName,
			DateOfBirth: emp.DateOfBirth,
			StartDate:   emp.StartDate,
			Department:  emp.Department.String(),
			Address: EmployeeAddress{
				Street:  emp.Street,
				City:    emp.City,
				State:   emp.State.String(),
				ZipCode: emp.ZipCode,
			},
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employeesDTO)
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
