package api

import (
	"encoding/json"
	"fmt"
)

type OriginateReq struct {
	rpBase
	Caller        string `json:"Caller" alias:"Caller"`
	Called        string `json:"DestNumber" alias:"DestNumber"`
	CalledDisplay string `json:"SpNumber" alias:"SpNumber"`
	AgentId       string `json:"AgentId" alias:"AgentId"`
	AppId         string `json:"app_id" alias:"app_id"`
}

func newOriginate(rpb rpBase, params map[string]interface{}) (retRp *OriginateReq, retErr error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("recover: %+v", r)
			retErr = fmt.Errorf("recover : %+v", r)
			retRp = nil
		}
	}()
	rp := OriginateReq{
		rpBase:        rpb,
		AgentId:       params["agent_id"].(string),
		CalledDisplay: params["called_display"].(string),
		Called:        params["called"].(string),
		Caller:        params["caller"].(string),
		AppId:         params["app_id"].(string),
	}
	return &rp, nil
}

// GetPayload ： 返回需要发送的字符串
func (rp *OriginateReq) GetPayload() ([]byte, error) {
	return json.Marshal(rp)
}
