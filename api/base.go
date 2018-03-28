package api

type rpBase struct {
	Ver       string `json:"ver,omitempty"`
	Ua        string `json:"ua,omitempty"`
	SeqID     int32  `json:"seq_id,omitempty"`
	MsgID     string `json:"msg_id,omitempty"`
	MsgType   string `json:"msg_type,omitempty"`
	Method    string `json:"method,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}
