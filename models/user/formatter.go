package user

type UserFormatted struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func UserFormatter(user User) UserFormatted {
	userFormatted := UserFormatted{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Username: user.Username,
	}

	return userFormatted
}
