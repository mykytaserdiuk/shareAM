package models

type User struct {
	ID    string  `db:"id"`
	Email *string `db:"email"`
}

type Credentials struct {
	Name     string `json:"name" db:"name"`
	Password string `json:"password"`
	Hash     string `db:"hash"`
}
