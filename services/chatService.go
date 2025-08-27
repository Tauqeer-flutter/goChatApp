package services

import (
	"goChatApp/domain"
	requests "goChatApp/handler/requests/chat"
)

type ChatService struct {
	Repository      domain.ChatRepositoryInterface
	GroupRepository domain.GroupRepositoryInterface
	userRepo        domain.UserRepositoryInterface
}

func (c *ChatService) SendMessage(request *requests.SendMessageRequest) (*domain.Chat, error) {
	if request.GroupId != nil {
		_, err := c.GroupRepository.GetGroupById(request.GroupId)
		if err != nil {
			return nil, err
		}
	} else {
		group := domain.Group{
			MemberCount: 2,
			GroupType:   "private",
		}
		groupId, err := c.GroupRepository.Create(&group)
		if err != nil {
			return nil, err
		}
		request.GroupId = &groupId
	}
	_, err := c.userRepo.GetById(request.SenderId)
	if err != nil {
		return nil, err
	}
	return c.Repository.CreateChat(request)
}

func (c *ChatService) List(groupId int64) ([]*domain.Chat, error) {
	return c.Repository.List(groupId)
}

func NewChatService(repository domain.ChatRepositoryInterface, groupRepo domain.GroupRepositoryInterface, userRepo domain.UserRepositoryInterface) *ChatService {
	return &ChatService{Repository: repository, GroupRepository: groupRepo, userRepo: userRepo}
}
