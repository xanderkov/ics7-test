package dto

type Patient struct {
	Id             int
	Surname        string
	Name           string
	Patronymic     string
	Height         int
	Weight         float64
	RoomNumber     int
	DegreeOfDanger int
}

type Patients []*Patient

type CreatePatient struct {
	Surname        string
	Name           string
	Patronymic     string
	Height         int
	Weight         float64
	RoomNumber     int
	DegreeOfDanger int
}

type UpdatePatient struct {
	Surname        string
	Name           string
	Patronymic     string
	Height         int
	Weight         float64
	RoomNumber     int
	DegreeOfDanger int
}
