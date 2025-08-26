package services

import (
	"fmt"
	"goChatApp/domain"
)

type GroupService struct {
	repository     domain.GroupRepositoryInterface
	userRepository domain.UserRepositoryInterface
}

func (g GroupService) Create(group *domain.Group, currentUserId *int64, otherUserId *int64) (*domain.Group, error) {
	if group.GroupType == "private" {
		otherUser, err := g.userRepository.GetById(*otherUserId)
		if err != nil {
			return nil, err
		}
		fmt.Println("Other user exists: ", otherUser)
	}
	groupId, err := g.repository.Create(group)
	otherMember := &domain.Member{
		MemberId: *otherUserId,
		GroupId:  groupId,
	}
	currentUser := &domain.Member{
		MemberId: *currentUserId,
		GroupId:  groupId,
	}
	err = g.repository.CreateMember(otherMember)
	if err != nil {
		return nil, err
	}
	err = g.repository.CreateMember(currentUser)
	if err != nil {
		return nil, err
	}
	fmt.Println("GROUP created: ", groupId)
	createdGroup, err := g.repository.GetGroupById(&groupId)
	if err != nil {
		return nil, err
	}
	return createdGroup, nil
}

func (g GroupService) List(userId int64) ([]*domain.Group, error) {
	return g.repository.List(userId)
}

func NewGroupService(repository domain.GroupRepositoryInterface, userRepository domain.UserRepositoryInterface) *GroupService {
	return &GroupService{
		repository:     repository,
		userRepository: userRepository,
	}
}
