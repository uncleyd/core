package e

var MsgFlags = map[int]string{
	MSG_OK:                         "ok",
	MSG_ERR:                        "fail",
	MSG_ERR_REQUEST:                "请求错误",
	MSG_ERR_SYNC:                   "签名错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "请求Token失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已过期，请重新获取",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[MSG_ERR]
}

func InitMsgFlags(msg map[int]string) {
	for k, v := range msg {
		MsgFlags[k] = v
	}
}
