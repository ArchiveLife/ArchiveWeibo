package model

type ReferenceType int

const (
	RefTypeAuthor ReferenceType = 1
)

type Reference struct {
	Type        ReferenceType
	ReferenceId string
}
