package repositories

import (
	"goChatApp/domain"
	"goChatApp/handler/requests/chat"

	"gorm.io/gorm"
)

type chatRepository struct {
	DB *gorm.DB
}

func (cr chatRepository) CreateChat(request *requests.SendMessageRequest) (*domain.Chat, error) {
	chat := &domain.Chat{
		GroupId:     *request.GroupId,
		Message:     request.Message,
		SenderId:    request.SenderId,
		FileUrl:     request.FileUrl,
		ReferenceTo: &request.References,
	}
	err := cr.DB.Create(chat).Error
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (cr chatRepository) List(groupId int64) ([]*domain.Chat, error) {
	var chats []*domain.Chat
	err := cr.DB.Raw("SELECT * FROM go_chat_app.chats WHERE group_id = ? ORDER BY created_at DESC;", groupId).Scan(&chats).Error
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func NewChatRepository(db *gorm.DB) domain.ChatRepositoryInterface {
	return &chatRepository{
		DB: db,
	}
}
