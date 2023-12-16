package dto

type UserResponse struct {
	UserID uint `json:"userId"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
