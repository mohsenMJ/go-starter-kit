package requests

type RegisterRequest struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:3"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
