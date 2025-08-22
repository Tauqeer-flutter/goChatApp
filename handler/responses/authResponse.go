package responses

import "goChatApp/domain"

type AuthResponse struct {
	AccessToken string      `json:"access_token"`
	User        domain.User `json:"user"`
}
