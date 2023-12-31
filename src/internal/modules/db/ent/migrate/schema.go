// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DiseasesColumns holds the columns for the "diseases" table.
	DiseasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "threat", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "degree_of_danger", Type: field.TypeInt},
	}
	// DiseasesTable holds the schema information for the "diseases" table.
	DiseasesTable = &schema.Table{
		Name:       "diseases",
		Columns:    DiseasesColumns,
		PrimaryKey: []*schema.Column{DiseasesColumns[0]},
	}
	// DoctorsColumns holds the columns for the "doctors" table.
	DoctorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token_id", Type: field.TypeString, Unique: true},
		{Name: "surname", Type: field.TypeString},
		{Name: "speciality", Type: field.TypeString},
		{Name: "role", Type: field.TypeString},
	}
	// DoctorsTable holds the schema information for the "doctors" table.
	DoctorsTable = &schema.Table{
		Name:       "doctors",
		Columns:    DoctorsColumns,
		PrimaryKey: []*schema.Column{DoctorsColumns[0]},
	}
	// PatientsColumns holds the columns for the "patients" table.
	PatientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "surname", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "patronymic", Type: field.TypeString},
		{Name: "height", Type: field.TypeInt},
		{Name: "weight", Type: field.TypeFloat64},
		{Name: "degree_of_danger", Type: field.TypeInt},
		{Name: "disease_has", Type: field.TypeInt, Nullable: true},
		{Name: "room_number", Type: field.TypeInt},
	}
	// PatientsTable holds the schema information for the "patients" table.
	PatientsTable = &schema.Table{
		Name:       "patients",
		Columns:    PatientsColumns,
		PrimaryKey: []*schema.Column{PatientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "patients_diseases_has",
				Columns:    []*schema.Column{PatientsColumns[7]},
				RefColumns: []*schema.Column{DiseasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "patients_rooms_contains",
				Columns:    []*schema.Column{PatientsColumns[8]},
				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number", Type: field.TypeInt, Unique: true},
		{Name: "floor", Type: field.TypeInt},
		{Name: "number_beds", Type: field.TypeInt},
		{Name: "number_patients", Type: field.TypeInt},
		{Name: "type_room", Type: field.TypeString},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:       "rooms",
		Columns:    RoomsColumns,
		PrimaryKey: []*schema.Column{RoomsColumns[0]},
	}
	// DoctorPatientColumns holds the columns for the "doctor_patient" table.
	DoctorPatientColumns = []*schema.Column{
		{Name: "doctor_id", Type: field.TypeInt},
		{Name: "patient_id", Type: field.TypeInt},
	}
	// DoctorPatientTable holds the schema information for the "doctor_patient" table.
	DoctorPatientTable = &schema.Table{
		Name:       "doctor_patient",
		Columns:    DoctorPatientColumns,
		PrimaryKey: []*schema.Column{DoctorPatientColumns[0], DoctorPatientColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "doctor_patient_doctor_id",
				Columns:    []*schema.Column{DoctorPatientColumns[0]},
				RefColumns: []*schema.Column{DoctorsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "doctor_patient_patient_id",
				Columns:    []*schema.Column{DoctorPatientColumns[1]},
				RefColumns: []*schema.Column{PatientsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DiseasesTable,
		DoctorsTable,
		PatientsTable,
		RoomsTable,
		DoctorPatientTable,
	}
)

func init() {
	PatientsTable.ForeignKeys[0].RefTable = DiseasesTable
	PatientsTable.ForeignKeys[1].RefTable = RoomsTable
	DoctorPatientTable.ForeignKeys[0].RefTable = DoctorsTable
	DoctorPatientTable.ForeignKeys[1].RefTable = PatientsTable
}
