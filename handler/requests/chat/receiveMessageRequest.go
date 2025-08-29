package requests

type ReceiveMessageRequest struct {
	Message string  `json:"message"`
	FileUrl *string `json:"file_url"`
}
