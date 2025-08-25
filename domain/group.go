package domain

import "time"

type Group struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	GroupType   string    `json:"group_type" gorm:"not null"` // Can be either private or group
	MemberCount int       `json:"member_count" gorm:"not null"`
	Members     []Member  `json:"members" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null"`
}

type Member struct {
	MemberId  int64     `json:"id"`
	GroupId   int64     `json:"group_id"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	User      User      `gorm:"foreignKey:MemberId" json:"-"`
}

type GroupRepositoryInterface interface {
	Create(group *Group) (int64, error)
	CreateMember(member *Member) error
	GetGroupById(groupId *int64) (*Group, error)
}

type GroupServiceInterface interface {
	Create(group *Group, currentUserId *int64, otherUserId *int64) (*Group, error)
}
