package dto

type Disease struct {
	Id             int
	Name           string
	Threat         string
	DegreeOfDanger int
}

type Diseases []*Disease

type CreateDisease struct {
	Name           string
	Threat         string
	DegreeOfDanger int
}

type UpdateDisease struct {
	Name           string
	Threat         string
	DegreeOfDanger int
}
