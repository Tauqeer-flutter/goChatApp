package repositories

import (
	"fmt"
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

func (g groupRepository) List(userId int64) ([]*domain.Group, error) {
	var groups []*domain.Group
	var members []domain.Member
	err := g.Db.Raw("SELECT m.*, g.id, g.name, g.description, g.group_type, g.member_count FROM go_chat_app.members m JOIN go_chat_app.groups g ON g.id = group_id WHERE m.member_id = ? ORDER BY m.updated_at DESC;", userId).Scan(&groups).Error
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		id := group.Id
		fmt.Println("id: ", id)
		err = g.Db.Table("members").Where("group_id = ?;", group.Id).Scan(&members).Error
		if err != nil {
			fmt.Println("Error fetching members for group: ", id, err)
			continue
		}
		group.MemberCount = len(members)
		group.Members = members
	}
	return groups, nil
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

func NewGroupRepository(db *gorm.DB) domain.GroupRepositoryInterface {
	return &groupRepository{
		Db: db,
	}
}
