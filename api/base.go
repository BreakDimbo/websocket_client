package api

type rpBase struct {
	Ver       string `json:"Ver,omitempty"`
	Ua        string `json:"Ua,omitempty"`
	SeqID     int32  `json:"SeqId,omitempty"`
	MsgID     string `json:"MsgId,omitempty"`
	MsgType   string `json:"MsgType,omitempty"`
	Method    string `json:"Method,omitempty"`
	Timestamp int64  `json:"Timestamp,omitempty"`
}
