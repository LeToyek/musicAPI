package entities

type User struct {
	ID            string `json:"id" `
	Username      string `json:"username" binding:"required,min=2,max=50"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required"`
	Refresh_Token string `json:"refresh_token"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
}
type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}
