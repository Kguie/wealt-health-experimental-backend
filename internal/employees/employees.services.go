package employees

import (
	"context"
	"fmt"
	"wealth-health-backend/ent"

	"github.com/google/uuid"
)

type EmployeeHandler struct {
	Client *ent.Client
}

// NewEmployeeHandler crée une instance de `EmployeeHandler`
func NewEmployeeHandler(client *ent.Client) *EmployeeHandler {
	return &EmployeeHandler{Client: client}
}

func (employee *EmployeeHandler) Create(ctx context.Context, dto CreateEmployeeDTO) (map[string]string, error) {
	_, err := employee.Client.Employee.
		Create().SetFirstName(dto.FirstName).
		SetLastName(dto.LastName).
		SetDateOfBirth(dto.DateOfBirth.Time).
		SetStartDate(dto.StartDate.Time).
		SetDepartment(dto.Department).
		SetStreet(dto.Address.Street).
		SetCity(dto.Address.City).
		SetState(dto.Address.State).
		SetZipCode(dto.Address.ZipCode).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création de l'employé: %w", err)
	}
	response := map[string]string{"message": "L'employé a été ajouté avec succès."}
	return response, nil
}

func (employee *EmployeeHandler) List(ctx context.Context) (*[]EmployeeDTO, error) {

	employees, err := employee.Client.Employee.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des employés: %w", err)
	}

	var employeesDTO []EmployeeDTO
	for _, emp := range employees {
		employeesDTO = append(employeesDTO, formatToEmployeeDTO(*emp))
	}

	return &employeesDTO, nil
}

func (employee *EmployeeHandler) GetByID(ctx context.Context, id uuid.UUID) (EmployeeDTO, error) {
	emp, err := employee.Client.Employee.Get(ctx, id)
	if err != nil {
		return EmployeeDTO{}, fmt.Errorf("erreur lors de la récupération de l'employé: %w", err)
	}
	res := formatToEmployeeDTO(*emp)
	return res, nil
}

func (employee *EmployeeHandler) UpdateByID(ctx context.Context, id uuid.UUID, dto UpdateEmployeeDTO) (EmployeeDTO, error) {
	query := employee.Client.Employee.UpdateOneID(id)
	if dto.DateOfBirth != nil {
		query.SetDateOfBirth(dto.DateOfBirth.Time)
	}
	if dto.StartDate != nil {
		query.SetStartDate(dto.StartDate.Time)
	}
	if dto.Department != nil {
		query.SetDepartment(*dto.Department)
	}
	if dto.Address != nil {
		if dto.Address.Street != nil {
			query.SetStreet(*dto.Address.Street)
		}
		if dto.Address.City != nil {
			query.SetCity(*dto.Address.City)
		}
		if dto.Address.State != nil {
			query.SetState(*dto.Address.State)
		}
		if dto.Address.ZipCode != nil {
			query.SetZipCode(*dto.Address.ZipCode)
		}
	}

	// Sauvegarder les modifications
	updatedEmp, err := query.Save(ctx)
	if err != nil {
		return EmployeeDTO{}, fmt.Errorf("❌ erreur lors de la mise à jour de l'employé: %w", err)
	}

	// Transformer en DTO avant de renvoyer la réponse
	res := formatToEmployeeDTO(*updatedEmp)
	return res, nil
}

func (employee *EmployeeHandler) DeleteByID(ctx context.Context, id uuid.UUID) (map[string]string, error) {
	err := employee.Client.Employee.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("❌ Erreur lors de la suppression:  %w", err)
	}
	response := map[string]string{"message": "L'employé a été supprimé avec succès."}
	return response, nil
}

func formatToEmployeeDTO(emp ent.Employee) EmployeeDTO {
	address := EmployeeAddress{
		Street:  emp.Street,
		City:    emp.City,
		State:   emp.State,
		ZipCode: emp.ZipCode,
	}
	employee := EmployeeDTO{
		Address:     address,
		Id:          emp.ID.String(),
		FirstName:   emp.FirstName,
		LastName:    emp.LastName,
		DateOfBirth: emp.DateOfBirth,
		StartDate:   emp.StartDate,
		Department:  emp.Department,
	}
	return employee

}
