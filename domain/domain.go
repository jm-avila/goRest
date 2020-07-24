package domain

// UserRepo methods to interface with the db
type UserRepo interface {
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(user *User) (*User, error)
}
// DB connection
type DB struct {
	UserRepo UserRepo
}

// Domain DB instance
type Domain struct {
	DB DB
}

