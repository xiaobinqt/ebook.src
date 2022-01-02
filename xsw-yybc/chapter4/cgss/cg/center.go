package cg

import (
	"encoding/json"
	"fmt"
	"sync"

	"go.src/xsw-yybc/chapter4/cgss/ipc"
)

var _ ipc.IpcServer = &CenterServer{} // 确认实现了 Server 接口

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

type CenterServer struct {
	servers map[string]ipc.Server
	Players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)

	return &CenterServer{
		servers: servers,
		Players: players,
	}
}

func (server *CenterServer) addPlayer(params string) (err error) {
	player := NewPlayer()
	err = json.Unmarshal([]byte(params), &player)
	if err != nil {
		fmt.Println("addPlayer unmarshal err", err.Error())
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	// 偷懒了,没做重复检查
	server.Players = append(server.Players, player)

	return nil
}

func (server *CenterServer) removePlayer(params string) (err error) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for i, v := range server.Players {
		if v.Name == params {
			if len(server.Players) == 1 {
				server.Players = make([]*Player, 0)
			} else if i == len(server.Players)-1 {
				server.Players = server.Players[1:]
			} else if i == 0 {
				server.Players = server.Players[1:]
			} else {
				server.Players = append(server.Players[:i-1],
					server.Players[:i+1]...)
			}
			return nil
		}
	}

	return fmt.Errorf("Player not found")
}

func (server *CenterServer) listPlayer(params string) (players string,
	err error) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if len(server.Players) > 0 {
		b, _ := json.Marshal(server.Players)
		players = string(b)
	} else {
		err = fmt.Errorf("No player online")
	}

	return
}

func (server *CenterServer) broadcast(params string) (err error) {
	var message Message
	err = json.Unmarshal([]byte(params), &message)
	if err != nil {
		fmt.Println("CenterServer broadcast err", err.Error())
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	if len(server.Players) > 0 {
		for _, player := range server.Players {
			player.mq <- &message
		}
	} else {
		err = fmt.Errorf("No player online")
	}

	return err
}

func (server *CenterServer) Handle(method, params string) *ipc.Response {
	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{
				Code: err.Error(),
			}
		}
	case "removeplayer":
		err := server.removePlayer(params)
		if err != nil {
			return &ipc.Response{
				Code: err.Error(),
			}
		}
	case "listplayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &ipc.Response{
				Code: err.Error(),
			}
		}
		return &ipc.Response{
			Code: "200",
			Body: players,
		}
	case "broadcast":
		err := server.broadcast(params)
		if err != nil {
			return &ipc.Response{
				Code: err.Error(),
			}
		}
		return &ipc.Response{
			Code: "200",
		}
	default:
		return &ipc.Response{
			Code: "404",
			Body: method + ":" + params,
		}
	}

	return &ipc.Response{Code: "200"}
}

func (server *CenterServer) Name() string {
	return "CenterServer"
}
