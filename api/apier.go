package api

import (
	"fmt"

	"github.com/satori/go.uuid"

	"github.com/breakD/wb_client/code"
	log "github.com/breakD/wb_client/logging"
	"github.com/breakD/wb_client/wbclient"
)

func Execute(cmd string, c *wbclient.Client) error {
	msgID := uuid.NewV4().String()
	switch cmd {
	case "login":
		send(c, code.REQ, code.LoginMethod, msgID, &map[string]interface{}{
			"agent_id": "1441",
			"app_id":   "f781fd92-9e96-4195-7a7e-257307edb4ad",
			"state":    "idle",
			"token":    "tmp:token",
		})
	}
	return nil
}

func send(c *wbclient.Client, typ string, method string, msgID string, custom *map[string]interface{}) error {
	rep, err := NewReport(&NewParam{
		SeqID:   int32(1),
		MsgID:   msgID,
		MsgType: typ,
		Method:  method,
		Custom:  *custom,
	})
	if err != nil {
		fmt.Errorf("new report error : %s", err)
		return err
	}
	payload, err := rep.GetPayload()
	if err != nil {
		fmt.Errorf("reporter.GetPayload error: %s", err)
		return err
	}

	// 向 ws 写数据
	err = c.SendMessage(payload)
	if err != nil {
		fmt.Errorf("ws write error: %s", err)
		return err
	}
	log.Debugf("send msg to sig: %s", payload)
	return nil
}
