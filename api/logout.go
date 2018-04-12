package api

import (
	"encoding/json"
	"fmt"
)

// Logout 登录请求
type Logout struct {
	rpBase
	AgentId string `json:"agent_id,omitempty"`
	AppID   string `json:"app_id,omitempty"`
}

func newLogout(rpb rpBase, params map[string]interface{}) (retRp *Logout, retErr error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("recover: %+v", r)
			retErr = fmt.Errorf("recover : %+v", r)
			retRp = nil
		}
	}()
	rp := Logout{
		rpBase:  rpb,
		AgentId: params["agent_id"].(string),
		AppID:   params["app_id"].(string),
	}
	return &rp, nil
}

// GetPayload ： 返回需要发送的字符串
func (rp *Logout) GetPayload() ([]byte, error) {
	return json.Marshal(rp)
}
