package repositories

import (
	"goChatApp/domain"
	"goChatApp/handler/requests"

	"gorm.io/gorm"
)

type chatRepository struct {
	DB *gorm.DB
}

func (cr chatRepository) CreateChat(request *requests.SendMessageRequest) error {
	chat := &domain.Chat{
		GroupId:     *request.GroupId,
		Message:     request.Message,
		SenderId:    request.SenderId,
		ReferenceTo: &request.References,
	}
	return cr.DB.Create(chat).Error
}

func NewChatRepository(db *gorm.DB) *chatRepository {
	return &chatRepository{
		DB: db,
	}
}
