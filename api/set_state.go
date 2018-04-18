package api

import (
	"encoding/json"
	"fmt"
)

// SetStateReq 登录请求
type SetStateReq struct {
	rpBase
	AgentId string `json:"AgentId,omitempty"`
	AppID   string `json:"AppId,omitempty"`
	State   string `json:"State,omitempty"`
	Token   string `json:"Token"`
}

func newSetState(rpb rpBase, params map[string]interface{}) (retRp *SetStateReq, retErr error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("recover: %+v", r)
			retErr = fmt.Errorf("recover : %+v", r)
			retRp = nil
		}
	}()
	rp := SetStateReq{
		rpBase:  rpb,
		AgentId: params["agent_id"].(string),
		AppID:   params["app_id"].(string),
		State:   params["state"].(string),
		Token:   params["token"].(string),
	}
	return &rp, nil
}

// GetPayload ： 返回需要发送的字符串
func (rp *SetStateReq) GetPayload() ([]byte, error) {
	return json.Marshal(rp)
}
