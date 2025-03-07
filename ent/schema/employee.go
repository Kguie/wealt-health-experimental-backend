package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Employee struct {
	ent.Schema
}

func (Employee) Fields() []ent.Field {

	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("firstName").NotEmpty(),
		field.String("lastName").NotEmpty(),
		field.Time("dateOfBirth"),
		field.Time("startDate"),
		field.String("department").NotEmpty(),
		field.String("street").NotEmpty(),
		field.String("city").NotEmpty(),
		field.String("state").NotEmpty(),
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
