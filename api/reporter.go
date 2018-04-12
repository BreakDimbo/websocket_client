package api

import (
	"fmt"

	"time"

	"github.com/breakD/wb_client/code"
)

// Reporter: 发送消息接口
type Reporter interface {
	GetPayload() ([]byte, error)
}

// NewParam 是NewReport时参数
type NewParam struct {
	SeqID   int32
	MsgID   string
	MsgType string
	Method  string
	Custom  map[string]interface{}
}

// NewReport ： 创建report消息的一个简单工厂方法
func NewReport(param *NewParam) (Reporter, error) {
	rpb := rpBase{
		Ver:       "1.0",
		Ua:        "robot-dan",
		SeqID:     param.SeqID,
		MsgID:     param.MsgID,
		MsgType:   param.MsgType,
		Method:    param.Method,
		Timestamp: time.Now().UnixNano() / 1000000,
	}

	msgType := param.MsgType
	methodStr := param.Method

	switch {
	case equal(code.REQ, code.LoginMethod, msgType, methodStr):
		return newLoginReq(rpb, param.Custom)
	case equal(code.REQ, code.LogoutMethod, msgType, methodStr):
		return newLogout(rpb, param.Custom)
	case equal(code.REQ, code.SetStateMethod, msgType, methodStr):
		return newSetState(rpb, param.Custom)
	case equal(code.REQ, code.DailoutMethod, msgType, methodStr):
		return newOriginate(rpb, param.Custom)
	}
	return nil, fmt.Errorf("unknow report method")
}

func equal(codeMsgType, codeMethod, msgType, methodStr string) bool {
	return codeMethod == methodStr && codeMsgType == msgType
}
