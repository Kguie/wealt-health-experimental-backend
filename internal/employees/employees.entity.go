package employees

import (
	"wealth-health-backend/pkg/data"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Employee struct {
	ent.Schema
}

func (Employee) Fields() []ent.Field {
	var departmentsValues []string
	var statesValues []string

	// Parcours de la tranche Departments pour extraire les valeurs
	for _, i := range data.Departments {
		departmentsValues = append(departmentsValues, i.Value)
	}

	for _, i := range data.States {
		statesValues = append(statesValues, i.Value)
	}

	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("firstName"),
		field.String("lastName"),
		field.Time("dateOfBirth"),
		field.Time("startDate"),
		field.Enum("department").Values(departmentsValues...),
		field.String("street"),
		field.String("city"),
		field.Enum("state").Values(statesValues...),
		field.String("zipCode"),
	}
}
