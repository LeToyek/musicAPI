package entities

type User struct {
	ID           string `json:"id" `
	Username     string `json:"username" validate:"required,min=2,max=50"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,password"`
	RefreshToken string `json:"refresh_token"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}
type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}
