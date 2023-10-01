package dto

// Сервис для аунтефикации в телеграмм

type Auth struct {
	Id      int
	TokenId string
}

type Users []*Auth

type NewDoctor struct {
	TokenId    string
	Surname    string
	Speciality string
	Role       string
}
