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
		Ua:        "stress-test",
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
		// case equal(code.REQ, code.GetChannelTokenMethod, msgType, methodStr):
		// 	return newGetChannelTokenReq(rpb, param.Custom)
		// case equal(code.REQ, code.InviteMethod, msgType, methodStr):
		// 	return newInviteReq(rpb, param.Custom)
		// case equal(code.ACK, code.InviteMethod, msgType, methodStr):
		// 	return newInviteAck(rpb, param.Custom)
		// case equal(code.REQ, code.AnswerMethod, msgType, methodStr):
		// 	return newAnswerReq(rpb, param.Custom)
		// case equal(code.ACK, code.AnswerMethod, msgType, methodStr):
		// 	return newAnswerAck(rpb, param.Custom)
		// case equal(code.REQ, code.ByeMethod, msgType, methodStr):
		// 	return newByeReq(rpb, param.Custom)
		// case equal(code.ACK, code.ByeMethod, msgType, methodStr):
		// 	return newByeAck(rpb, param.Custom)
		// case equal(code.REQ, code.CancelMethod, msgType, methodStr):
		// 	return newCancelReq(rpb, param.Custom)
		// case equal(code.ACK, code.CancelMethod, msgType, methodStr):
		// 	return newCancelAck(rpb, param.Custom)
		// case equal(code.NTF, code.StartMediaMethod, msgType, methodStr):
		// 	return newStartMediaNtf(rpb, param.Custom)
		// case equal(code.NTF, code.StopMediaMethod, msgType, methodStr):
		// 	return newStopMediaNtf(rpb, param.Custom)
	}
	return nil, fmt.Errorf("unknow report method")
}

func equal(codeMsgType, codeMethod, msgType, methodStr string) bool {
	return codeMethod == methodStr && codeMsgType == msgType
}
