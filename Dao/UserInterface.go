package Dao

type User interface {
	GetID() string
	SetID(string)
	GetPASS() string
	SetPASS(string)
}
