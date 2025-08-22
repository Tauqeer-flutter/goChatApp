package services

import (
	"errors"
	"goChatApp/domain"
)

type GroupService struct {
	repository     domain.GroupRepositoryInterface
	userRepository domain.UserRepositoryInterface
}

func (g GroupService) Create(group *domain.Group) (*domain.Group, error) {
	return nil, errors.New("Hehe")
}

func NewGroupService(repository domain.GroupRepositoryInterface, userRepository domain.UserRepositoryInterface) *GroupService {
	return &GroupService{
		repository:     repository,
		userRepository: userRepository,
	}
}
