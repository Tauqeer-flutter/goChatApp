package responses

type UserResponse struct {
	Id        int64   `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	PhotoUrl  *string `json:"photo_url"`
	Phone     *string `json:"phone"`
}
