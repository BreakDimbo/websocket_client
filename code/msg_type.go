package code

// 消息类型枚举值
const (
	REQ = "REQ" // 请求消息
	REP = "REP" // 相应消息
	EVT = "EVT" // 事件消息
	ACK = "ACK" // 事件相应
	NTF = "NTF" // 通知消息
)

var MsgTypeList = []string{
	REQ,
	REP,
	EVT,
	ACK,
	NTF,
}
