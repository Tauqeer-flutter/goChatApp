package requests

type ChatWebsocketRequest struct {
	GroupId int64 `json:"group_id" form:"group_id" binding:"required"`
}
