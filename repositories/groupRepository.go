package repositories

import (
	"goChatApp/domain"

	"gorm.io/gorm"
)

type groupRepository struct {
	Db *gorm.DB
}

func (g groupRepository) Create(group *domain.Group) (int64, error) {
	var newGroup domain.Group
	err := g.Db.Create(group).Scan(&newGroup).Error
	if err != nil {
		return -1, err
	}
	return newGroup.Id, nil
}

func (g groupRepository) CreateMember(member *domain.Member) error {
	return g.Db.Create(member).Error
}

func (g groupRepository) GetGroupById(groupId *int64) (*domain.Group, error) {
	var group domain.Group
	err := g.Db.Raw("SELECT * FROM go_chat_app.groups WHERE id = ?", groupId).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func NewGroupRepository(db *gorm.DB) *groupRepository {
	return &groupRepository{
		Db: db,
	}
}
