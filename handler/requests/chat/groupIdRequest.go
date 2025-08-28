package requests

type GroupIdRequest struct {
	GroupId int64 `json:"group_id" form:"group_id" binding:"required"`
}
