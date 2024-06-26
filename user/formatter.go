package user

type UserFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Occupation  string `json:"occupation"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Token       string `json:"token"`
	ImageURL    string `json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		Name:        user.Name,
		Occupation:  user.Occupation,
		Description: user.Description,
		Email:       user.Email,
		Token:       token,
		ImageURL:    user.AvatarFileName,
	}
	return formatter
}