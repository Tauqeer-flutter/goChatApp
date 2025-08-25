package domain

import (
	"goChatApp/handler/requests"
	"time"
)

type Chat struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	GroupId     int64     `json:"group_id"`
	Message     string    `json:"message"`
	SenderId    int64     `json:"sender_id"`
	ReferenceTo *int64    `json:"reference_to"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ChatRepositoryInterface interface {
	CreateChat(request *requests.SendMessageRequest) error
}

type ChatServiceInterface interface {
	SendMessage(request *requests.SendMessageRequest) error
}
