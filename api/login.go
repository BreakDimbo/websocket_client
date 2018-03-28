package api

import (
	"encoding/json"
	"fmt"
)

// LoginReq 登录请求
type LoginReq struct {
	rpBase
	AgentId string `json:"agent_id,omitempty"`
	AppID   string `json:"app_id,omitempty"`
	State   string `json:"state,omitempty"`
	Token   string `json:"token"`
}

func newLoginReq(rpb rpBase, params map[string]interface{}) (retRp *LoginReq, retErr error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("recover: %+v", r)
			retErr = fmt.Errorf("recover : %+v", r)
			retRp = nil
		}
	}()
	rp := LoginReq{
		rpBase:  rpb,
		AgentId: params["agent_id"].(string),
		AppID:   params["app_id"].(string),
		State:   params["state"].(string),
		Token:   params["token"].(string),
	}
	return &rp, nil
}

// GetPayload ： 返回需要发送的字符串
func (rp *LoginReq) GetPayload() ([]byte, error) {
	return json.Marshal(rp)
}
