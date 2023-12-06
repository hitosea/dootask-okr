package service

import (
	"dootask-okr/app/interfaces"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

var WebSocketService = webSocketService{}

type webSocketService struct{}

// 推送给指定fd类型全体用户
func (ws webSocketService) PushMsg(fd int, msg interface{}) {
	log.Println("推送任务所有连接列表:", interfaces.WsClients)
	for _, v := range interfaces.WsClients {
		if v.Rid == int32(fd) {
			log.Println("推送消息给fd:", v.Rid)
			log.Printf("推送消息内容:%v", msg)
			msgJSON, err := json.Marshal(msg)
			if err != nil {
				log.Println("Failed to convert message to JSON:", err)
				continue
			}
			err = v.Conn.WriteMessage(websocket.TextMessage, msgJSON)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// 推送给指定用户
func (ws webSocketService) PushMsgToUser(userid []int, msg interface{}) {
	log.Println("推送任务所有连接列表:", interfaces.WsClients)
	for _, v := range interfaces.WsClients {
		for _, uid := range userid {
			if v.Uid == int32(uid) {
				log.Println("推送消息给fd:", v.Rid)
				log.Printf("推送消息内容:%v", msg)
				msgJSON, err := json.Marshal(msg)
				if err != nil {
					log.Println("Failed to convert message to JSON:", err)
					continue
				}
				err = v.Conn.WriteMessage(websocket.TextMessage, msgJSON)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
