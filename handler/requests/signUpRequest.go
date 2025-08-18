package requests

type SignUpRequest struct {
	FirstName string  `json:"first_name" binding:"required"`
	LastName  string  `json:"last_name" binding:"required"`
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"password" binding:"required,min=7"`
	Phone     *string `json:"phone" binding:"omitempty,min=9,max=12"`
}
