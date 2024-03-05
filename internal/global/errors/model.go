package errors

const OK int = 200

const SERVER_COMMON_ERROR int = 50000
const DB_DO_ERROR int = 50100
const TOKEN_EXPIRE_ERROR int = 40300
const REQUEST_COMMON_ERROR int = 40400

var (
	SUCCESS         = newError(OK, "Success")
	REQUEST_ERROR   = newError(REQUEST_COMMON_ERROR, "无效的请求")
	DB_ERROR        = newError(DB_DO_ERROR, "数据库操作失败")
	AUTHORIZE_ERROR = newError(TOKEN_EXPIRE_ERROR, "鉴权失败")
	SERVER_ERROR    = newError(SERVER_COMMON_ERROR, "服务器内部错误")
)
