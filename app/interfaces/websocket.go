package interfaces

import "github.com/gorilla/websocket"

type WebsocketReq struct {
	Token    string `json:"token"`
	Language string `json:"language"`
}

type WsClient struct {
	Conn *websocket.Conn `json:"conn"`

	Type string `json:"type"` // 客户端类型（如：user）
	Uid  int32  `json:"uid"`  // 客户端用户ID（会员ID）
	Rid  int32  `json:"rid"`  // 客户端序号ID（WebSocket ID）
}

type WsMsg struct {
	Action int `json:"action"` // 消息类型：1、上线；2、下线；3、消息
	Data   any `json:"data"`   // 消息内容

	Type string `json:"type"` // 客户端类型（如：user）
	Uid  int32  `json:"uid"`  // 客户端用户ID（会员ID）
	Rid  int32  `json:"rid"`  // 客户端序号ID（WebSocket ID）
}

const (
	WsHeartbeat = 0 // 心跳
	WsOnline    = 1 // 连接
	WsOffline   = 2 // 断开
	WsSendMsg   = 3 // 消息发送

	WsIsUser    = "user"    // 会员
	WsIsUnknown = "unknown" // 未知
)

var (
	WsClients []WsClient
)
