package api

import (
	"fmt"
	"strings"

	"github.com/satori/go.uuid"

	"github.com/breakD/wb_client/code"
	log "github.com/breakD/wb_client/logging"
	"github.com/breakD/wb_client/wbclient"
)

func Execute(cmd string, c *wbclient.Client) error {
	msgID := uuid.NewV4().String()
	switch strings.TrimSpace(cmd) {
	case code.LoginMethod:
		send(c, code.REQ, code.LoginMethod, msgID, &map[string]interface{}{
			"agent_id": "1443",
			"app_id":   "f781fd92-9e96-4195-7a7e-257307edb4ad",
			"state":    "idle",
			"token":    "tmp:token",
		})
	case "set_state_idle":
		send(c, code.REQ, code.SetStateMethod, msgID, &map[string]interface{}{
			"agent_id": "1443",
			"app_id":   "f781fd92-9e96-4195-7a7e-257307edb4ad",
			"state":    "idle",
			"token":    "tmp_token",
		})
	case "set_state_busy":
		send(c, code.REQ, code.SetStateMethod, msgID, &map[string]interface{}{
			"agent_id": "1443",
			"app_id":   "f781fd92-9e96-4195-7a7e-257307edb4ad",
			"state":    "busy",
			"token":    "tmp_token",
		})

	case "dialout":
		send(c, code.REQ, code.DailoutMethod, msgID, &map[string]interface{}{
			"agent_id":       "1443",
			"app_id":         "f781fd92-9e96-4195-7a7e-257307edb4ad",
			"called_display": "057126200670",
			"caller":         "92795089691002",
			"called":         "13240345451",
		})
	default:
		fmt.Println("[Error] wrong api command not match")
	}
	return nil
}

func send(c *wbclient.Client, typ string, method string, msgID string, custom *map[string]interface{}) (err error) {
	rep, err := NewReport(&NewParam{
		SeqID:   int32(1),
		MsgID:   msgID,
		MsgType: typ,
		Method:  method,
		Custom:  *custom,
	})
	if err != nil {
		fmt.Printf("[Error] new report error : %s\n", err)
		return
	}
	payload, err := rep.GetPayload()
	if err != nil {
		fmt.Printf("reporter.GetPayload error: %s\n", err)
		return
	}

	// 向 ws 写数据
	err = c.SendMessage(payload)
	if err != nil {
		fmt.Printf("ws write error: %s\n", err)
		return
	}

	log.Debugf("send msg to server: %s\n", payload)
	return
}
