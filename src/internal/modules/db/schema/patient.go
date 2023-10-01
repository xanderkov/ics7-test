package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Patient holds the schema definition for the Patient entity.
type Patient struct {
	ent.Schema
}

// Fields of the Patient.
func (Patient) Fields() []ent.Field {
	return []ent.Field{
		field.String("surname"),
		field.String("name"),
		field.String("patronymic"),
		field.Int("height"),
		field.Float("weight"),
		field.Int("roomNumber"),
		field.Int("degreeOfDanger"),
	}
}

// Edges of the Patient.
func (Patient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("repo", Room.Type).
			Ref("contains").
			Field("roomNumber").
			Unique().
			Required(),
		edge.From("doctor", Doctor.Type).
			Ref("treats"),
		edge.From("ills", Disease.Type).
			Ref("has").
			Unique(),
	}
}
