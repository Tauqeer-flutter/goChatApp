package responses

import "goChatApp/domain"

type SuccessAuthResponse struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	AccessToken string      `json:"access_token"`
	User        domain.User `json:"user"`
}
