package domain

import (
	"mime/multipart"
)

type MediaServiceInterface interface {
	UploadChatFile(groupId int64, file *multipart.FileHeader) (*string, error)
}
