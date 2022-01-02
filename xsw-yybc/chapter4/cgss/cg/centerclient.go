package cg

import (
	"encoding/json"
	"fmt"

	"go.src/xsw-yybc/chapter4/cgss/ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) (err error) {
	b, err := json.Marshal(*player)
	if err != nil {
		fmt.Println("CenterClient AddPlayer marshal err", err.Error())
		return err
	}

	resp, err := client.Call("addplayer", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}

	fmt.Println("CenterClient AddPlayer call err", err.Error())
	return err
}

func (client *CenterClient) RemovePlayer(name string) error {
	ret, _ := client.Call("removeplayer", name)
	if ret.Code == "200" {
		return nil
	}

	return fmt.Errorf("CenterClient RemovePlayer call err", ret.Code)
}

func (client *CenterClient) ListPlayer(params string) (ps []*Player,
	err error) {
	ps = make([]*Player, 0)
	resp, _ := client.Call("listplayer", params)
	if resp.Code != "200" {
		err = fmt.Errorf("CenterClient ListPlayer call err", resp.Code)
		return
	}

	err = json.Unmarshal([]byte(resp.Body), &ps)
	return
}

func (client *CenterClient) Broadcast(message string) error {
	m := &Message{
		Content: message,
	}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("CenterClient Broadcast marshal err", err.Error())
		return err
	}

	resp, _ := client.Call("broadcast", string(b))
	if resp.Code == "200" {
		return nil
	}

	return fmt.Errorf("CenterClient Broadcast %s ", resp.Code)

}
