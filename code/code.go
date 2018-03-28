package code

const (
	Ver             = "1.0"
	OK              = "000000" // 成功
	ParamsInvalid   = "000002" // 参数错误
	WSRedisError    = "000003" //
	TimeoutError    = "000035" // rest api time out
	InternalError   = "000036" // other errors
	LoginAgainError = "000040"
	LoginTokenError = "000041" //

	CalChannelTokenError      = "000100" // 计算 channelToken 错误
	ClientNotLoginError       = "000101" // 客户端为登录错误
	ClientTargetStateWrong    = "000102" // 设置坐席目标状态错误
	ClientNotInIdleStateError = "000103" // 目标坐席非空闲
)

var msg map[string]string

// Msg 用于返回对应 code 的文本信息
func Msg(code string) string {
	return msg[code]
}

// Init 用于初始化错误信息映射表
func Init() {
	msg = map[string]string{
		OK:              "成功",
		ParamsInvalid:   "参数错误",
		TimeoutError:    "接口响应超时", // rest api time out
		InternalError:   "内部错误",   // other errors
		WSRedisError:    "WS Redis 操作错误",
		LoginAgainError: "您已经完成登入",
		LoginTokenError: "登入token不合法", // 不合法或长度不对

		CalChannelTokenError:      "计算channel_token错误",
		ClientNotLoginError:       "用户未登录错误",
		ClientTargetStateWrong:    "设置坐席目标状态错误（只允许为空闲或忙碌）",
		ClientNotInIdleStateError: "目标坐席状态非空闲",
	}
}

// Cancel Evt 的 reason
const (
	CalledReject  = "001"
	CallerHangup  = "002"
	TimeoutCancel = "003"
)
