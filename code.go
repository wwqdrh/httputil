package httputil

type ResponseCode uint32

const (
	ServerOK ResponseCode = iota
)

const (
	// 参数问题
	ParamEmpty ResponseCode = 101 + iota
	ParamInvalid
)

func (c ResponseCode) String() string {
	switch c {
	case ServerOK:
		return "正常"
	case ParamEmpty:
		return "参数为空"
	case ParamInvalid:
		return "参数非法"
	default:
		return "未知状态码"
	}
}
