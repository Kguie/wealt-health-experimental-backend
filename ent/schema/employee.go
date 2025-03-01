package schema

import (
	"wealth-health-backend/pkg/data"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.UUID("id", uuid.UUID{}).Immutable(),
		field.String("firstName").NotEmpty(),
		field.String("lastName").NotEmpty(),
		field.Time("dateOfBirth"),
		field.Time("startDate"),
		field.Enum("department").Values(departmentsValues...),
		field.String("street").NotEmpty(),
		field.String("city").NotEmpty(),
		field.Enum("state").Values(statesValues...),
		field.String("zipCode").NotEmpty(),
	}
}

// Indexes
func (Employee) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("firstName", "lastName").
			Unique(),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return nil
}
