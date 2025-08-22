package domain

import "time"

type Group struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	GroupType   string    `json:"group_type" gorm:"not null"` // Can be either private or group
	MemberCount int64     `json:"member_count" gorm:"not null"`
	Members     []Member  `json:"members" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null"`
}

type Member struct {
	MemberId string `json:"id"`
	GroupId  string `json:"group_id"`
	User     User   `gorm:"foreignKey:MemberId" json:"-"`
}

type GroupRepositoryInterface interface {
	Create(group *Group) error
}

type GroupServiceInterface interface {
	Create(group *Group) (*Group, error)
}
