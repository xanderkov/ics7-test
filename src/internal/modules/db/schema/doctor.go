package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

//

type Doctor struct {
	ent.Schema
}

func (Doctor) Fields() []ent.Field {
	return []ent.Field{
		field.String("tokenId").Unique(),
		field.String("surname"),
		field.String("speciality"),
		field.String("role"),
	}
}

func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("treats", Patient.Type).StorageKey(
			edge.Table("doctor_patient"), edge.Columns("doctor_id", "patient_id"),
		),
	}
}
