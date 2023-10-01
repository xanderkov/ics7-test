// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"hospital/internal/modules/db/ent/room"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Room is the model entity for the Room schema.
type Room struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Number holds the value of the "number" field.
	Number int `json:"number,omitempty"`
	// Floor holds the value of the "floor" field.
	Floor int `json:"floor,omitempty"`
	// NumberBeds holds the value of the "numberBeds" field.
	NumberBeds int `json:"numberBeds,omitempty"`
	// NumberPatients holds the value of the "numberPatients" field.
	NumberPatients int `json:"numberPatients,omitempty"`
	// TypeRoom holds the value of the "typeRoom" field.
	TypeRoom string `json:"typeRoom,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomQuery when eager-loading is set.
	Edges        RoomEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoomEdges holds the relations/edges for other nodes in the graph.
type RoomEdges struct {
	// Contains holds the value of the contains edge.
	Contains []*Patient `json:"contains,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ContainsOrErr returns the Contains value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) ContainsOrErr() ([]*Patient, error) {
	if e.loadedTypes[0] {
		return e.Contains, nil
	}
	return nil, &NotLoadedError{edge: "contains"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Room) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case room.FieldID, room.FieldNumber, room.FieldFloor, room.FieldNumberBeds, room.FieldNumberPatients:
			values[i] = new(sql.NullInt64)
		case room.FieldTypeRoom:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Room fields.
func (r *Room) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case room.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case room.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				r.Number = int(value.Int64)
			}
		case room.FieldFloor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field floor", values[i])
			} else if value.Valid {
				r.Floor = int(value.Int64)
			}
		case room.FieldNumberBeds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field numberBeds", values[i])
			} else if value.Valid {
				r.NumberBeds = int(value.Int64)
			}
		case room.FieldNumberPatients:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field numberPatients", values[i])
			} else if value.Valid {
				r.NumberPatients = int(value.Int64)
			}
		case room.FieldTypeRoom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field typeRoom", values[i])
			} else if value.Valid {
				r.TypeRoom = value.String
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Room.
// This includes values selected through modifiers, order, etc.
func (r *Room) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryContains queries the "contains" edge of the Room entity.
func (r *Room) QueryContains() *PatientQuery {
	return NewRoomClient(r.config).QueryContains(r)
}

// Update returns a builder for updating this Room.
// Note that you need to call Room.Unwrap() before calling this method if this Room
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Room) Update() *RoomUpdateOne {
	return NewRoomClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Room entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Room) Unwrap() *Room {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Room is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Room) String() string {
	var builder strings.Builder
	builder.WriteString("Room(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", r.Number))
	builder.WriteString(", ")
	builder.WriteString("floor=")
	builder.WriteString(fmt.Sprintf("%v", r.Floor))
	builder.WriteString(", ")
	builder.WriteString("numberBeds=")
	builder.WriteString(fmt.Sprintf("%v", r.NumberBeds))
	builder.WriteString(", ")
	builder.WriteString("numberPatients=")
	builder.WriteString(fmt.Sprintf("%v", r.NumberPatients))
	builder.WriteString(", ")
	builder.WriteString("typeRoom=")
	builder.WriteString(r.TypeRoom)
	builder.WriteByte(')')
	return builder.String()
}

// Rooms is a parsable slice of Room.
type Rooms []*Room
