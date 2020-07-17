package domain

// UserRepo methods to interface with the db
type UserRepo interface {
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, interface{})
	Create(data *User) (*User, error)
}

// DB connection
type DB struct {
	UserRepo UserRepo
}

type Domain struct {
	DB DB
}
