package res

type ErrorCode int

const (
	SettingsError           ErrorCode = 1001
	ArgumentError           ErrorCode = 1002 // 参数错误
	ErrorTooManyConnections           = 1008 // 连接数已满
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError:           "系统错误",
		ArgumentError:           "参数错误",
		ErrorTooManyConnections: "聊天室已满",
	}
)
