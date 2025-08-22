package repositories

import (
	"goChatApp/domain"
	"gorm.io/gorm"
)

type groupRepository struct {
	Db *gorm.DB
}

func (g groupRepository) Create(group *domain.Group) error {
	//TODO implement me
	panic("implement me")
}

func NewGroupRepository(db *gorm.DB) *groupRepository {
	return &groupRepository{
		Db: db,
	}
}
