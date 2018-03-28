package code

// 方法枚举值
const (
	GetAttrMethod         = "get_attr"
	GetAttrAllMethod      = "get_attr_all"
	GetChannelTokenMethod = "get_channel_token"
	GetStateMethod        = "get_state"
	GetUserStateMethod    = "get_user_state"
	LoginMethod           = "login"
	LogoutMethod          = "logout"
	SetAttrMethod         = "set_attr"
	SetStateMethod        = "set_state"
	StartMediaMethod      = "start_media"
	StopMediaMethod       = "stop_media"
	InviteMethod          = "invite"
	AnswerMethod          = "answer"
	ByeMethod             = "bye"
	CancelMethod          = "cancel"

	DisconnectMethod  = "disconnect"
	StartRecordMethod = "start_record"
	StopRecordMethod  = "stop_record"
)

var MethodList = []string{
	GetAttrMethod,
	GetAttrAllMethod,
	GetChannelTokenMethod,
	GetStateMethod,
	GetUserStateMethod,
	LoginMethod,
	LogoutMethod,
	SetAttrMethod,
	SetStateMethod,
	StartMediaMethod,
	StopMediaMethod,
	InviteMethod,
	AnswerMethod,
	ByeMethod,
	CancelMethod,
	DisconnectMethod,
	StartRecordMethod,
	StopRecordMethod,
}
