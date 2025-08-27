package domain

import (
	requests "goChatApp/handler/requests/chat"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Chat struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	GroupId     int64     `json:"group_id"`
	Message     string    `json:"message"`
	SenderId    int64     `json:"sender_id"`
	ReferenceTo *int64    `json:"reference_to"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var Upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
var Mutex sync.Mutex
var Clients = make(map[*Client]int64)

type Client struct {
	UserId int64
	Active bool
	Conn   *websocket.Conn
}

type ChatRepositoryInterface interface {
	CreateChat(request *requests.SendMessageRequest) (*Chat, error)
	List(groupId int64) ([]*Chat, error)
}

type ChatServiceInterface interface {
	SendMessage(request *requests.SendMessageRequest) (*Chat, error)
	List(groupId int64) ([]*Chat, error)
}
