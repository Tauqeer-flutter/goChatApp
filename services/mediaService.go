package services

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

type MediaService struct{}

func (m *MediaService) UploadChatFile(groupId int64, file *multipart.FileHeader) (*string, error) {
	splits := strings.Split(file.Filename, ".")
	if len(splits) < 2 {
		return nil, fmt.Errorf("invalid file, filename should have extension")
	}
	extension := splits[len(splits)-1]
	filename := fmt.Sprintf("%d.%s", time.Now().Unix(), extension)
	path := "uploads/chats/" + fmt.Sprintf("%d", groupId) + "/"
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}
	bytes := make([]byte, file.Size)
	openedFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer openedFile.Close()
	if _, err := openedFile.Read(bytes); err != nil {
		return nil, err
	}
	if err := os.WriteFile(path+filename, bytes, os.ModePerm); err != nil {
		return nil, err
	}
	return &filename, nil
}

func NewMediaService() *MediaService {
	return &MediaService{}
}
