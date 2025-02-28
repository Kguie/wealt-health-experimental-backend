package employees

import (
	"fmt"
	"net/http"
)

func (employee *Employee) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an Employee")
}

func (employee *Employee) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all Employees")
}

func (employee *Employee) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an Employee by ID")
}

func (employee *Employee) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an Employee by ID")
}

func (employee *Employee) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an Employee by ID")
}
