package requests

type ReceiveMessageRequest struct {
	GroupId int64   `json:"group_id"`
	Message string  `json:"message"`
	FileUrl *string `json:"file_url"`
}
