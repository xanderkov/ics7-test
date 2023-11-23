package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Disease holds the schema definition for the Disease entity.
type Disease struct {
	ent.Schema
}

// Fields of the Disease.
func (Disease) Fields() []ent.Field {
	return []ent.Field{
		field.String("threat"),
		field.String("name"),
		field.Int("degreeOfDanger"),
	}
}

// Edges of the Disease.
func (Disease) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("has", Patient.Type),
	}
}
