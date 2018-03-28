package api

import (
	"encoding/json"
	"fmt"
)

// StopMediaNtf 登录请求
type StopMediaNtf struct {
	rpBase
	UserID    string `json:"user_id"`
	ChannelID string `json:"channel_id"`
}

func newStopMediaNtf(rpb rpBase, params map[string]interface{}) (retRp *StopMediaNtf, retErr error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("recover: %+v", r)
			retErr = fmt.Errorf("recover : %+v", r)
			retRp = nil
		}
	}()
	rp := StopMediaNtf{
		rpBase:    rpb,
		UserID:    params["user_id"].(string),
		ChannelID: params["channel_id"].(string),
	}
	return &rp, nil
}

// GetPayload ： 返回需要发送的字符串
func (rp *StopMediaNtf) GetPayload() ([]byte, error) {
	return json.Marshal(rp)
}
