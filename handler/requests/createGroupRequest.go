package requests

type CreateGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	GroupType   string `json:"group_type" binding:"required,min=3"`
	OtherUserId *int64 `json:"other_user_id"`
}
