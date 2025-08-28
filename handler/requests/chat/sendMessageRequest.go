package requests

type SendMessageRequest struct {
	GroupId    *int64  `json:"group_id"`
	Message    string  `json:"message" binding:"required"`
	SenderId   int64   `json:"sender_id" binding:"required"`
	FileUrl    *string `json:"file_url"`
	References int64   `json:"references"`
}
