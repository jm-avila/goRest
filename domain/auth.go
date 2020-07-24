package domain

// RegisterPayload request payload used for authentication
type RegisterPayload struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// Register to the app
func (d *Domain) Register(payload RegisterPayload) (*User, error) {
	userExists, _ := d.DB.UserRepo.GetByEmail(payload.Email)
	if userExists != nil {
		return nil, ErrUserWithEmailAlreadyExist
	}

	userExists, _ = d.DB.UserRepo.GetByUsername(payload.Username)
	if userExists != nil {
		return nil, ErrUserWithUsernameAlreadyExist
	}

	password, err := d.setPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	data := &User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: *password,
	}

	user, err := d.DB.UserRepo.Create(data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *Domain) setPassword(password string) (*string, error) {
	return nil, nil
}
