package services

import (
	"fmt"
	"goChatApp/domain"
	"goChatApp/handler/requests"
)

type ChatService struct {
	Repository      domain.ChatRepositoryInterface
	GroupRepository domain.GroupRepositoryInterface
	userRepo        domain.UserRepositoryInterface
}

func (c *ChatService) SendMessage(request *requests.SendMessageRequest) error {
	if request.GroupId != nil {
		_, err := c.GroupRepository.GetGroupById(request.GroupId)
		if err != nil {
			return err
		}
	} else {
		group := domain.Group{
			MemberCount: 2,
			GroupType:   "private",
		}
		groupId, err := c.GroupRepository.Create(&group)
		if err != nil {
			return err
		}
		request.GroupId = &groupId
	}
	user, err := c.userRepo.GetById(request.SenderId)
	if err != nil {
		return err
	}
	fmt.Println("Sender: ", user.Id, user.FirstName)
	return c.Repository.CreateChat(request)
}

func NewChatService(repository domain.ChatRepositoryInterface, groupRepo domain.GroupRepositoryInterface, userRepo domain.UserRepositoryInterface) *ChatService {
	return &ChatService{Repository: repository, GroupRepository: groupRepo, userRepo: userRepo}
}
