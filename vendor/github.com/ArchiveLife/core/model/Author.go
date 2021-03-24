package model

type Author struct {
	ID            ID
	Age           *int
	FullName      string
	Address       *string
	Mobile        *string
	ExtAttributes map[string]interface{}
}
