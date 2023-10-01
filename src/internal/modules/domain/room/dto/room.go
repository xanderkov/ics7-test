package dto

type Room struct {
	Id             int
	Num            int
	Floor          int
	NumberBeds     int
	TypeRoom       string
	NumberPatients int
}

type Rooms []*Room

type CreateRoom struct {
	Num            int
	Floor          int
	NumberBeds     int
	TypeRoom       string
	NumberPatients int
}

type UpdateRoom struct {
	Num            int
	Floor          int
	NumberBeds     int
	TypeRoom       string
	NumberPatients int
}
