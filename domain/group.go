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
	FirstName string    `json:"first_name" gorm:"->;column:first_name;-:migration"`
	LastName  string    `json:"last_name" gorm:"->;column:last_name;-:migration"`
	Email     string    `json:"email" gorm:"->;column:email;-:migration"`
	PhotoUrl  *string   `json:"photo_url" gorm:"->;column:photo_url;-:migration"`
	Phone     *string   `json:"phone" gorm:"->;column:phone;-:migration"`
}

//type

type GroupRepositoryInterface interface {
	Create(group *Group) (int64, error)
	List(userId int64) ([]*Group, error)
	CreateMember(member *Member) error
	GetGroupById(groupId *int64) (*Group, error)
}

type GroupServiceInterface interface {
	Create(group *Group, currentUserId *int64, otherUserId *int64) (*Group, error)
	List(currentUserId int64) ([]*Group, error)
}
