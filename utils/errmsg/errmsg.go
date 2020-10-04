// 错误信息
package errmsg

//import ""

/*
// 架构 框架
const (
	
)

var codemsg = map[int]string{
	
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
*/

const (
	SUCCSE = 200
	ERROR = 500

	// code= 1000... 用户模块的错误
	ERROR_USERNAME_USED = 1001 // 用户名重复
	ERROR_PASSWORD_WRONG = 1002 // 密码错误
	ERROR_USER_NOT_EXIST = 1003 // 用户名不能为空
	ERROR_TOKEN_EXIST = 1004 // 验证错误
	ERROR_TOKEN_RUNTIME = 1005 // 超时
	ERROR_TOKEN_WRONG = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT = 1008
	// code= 2000... 文章模块错误
	ERROR_CATENAME_USED = 2001 // 分类已存在
	ERROR_CATE_NOT_EXIST = 2002

	// code= 3000... 分类模块的错误

)

var codeMsg = map[int]string{
	SUCCSE: "OK",
	ERROR: "FAIL",
	ERROR_USERNAME_USED: "用户名已存在！",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_TOKEN_EXIST: "token不存在",
	ERROR_TOKEN_RUNTIME: "token已过期",
	ERROR_TOKEN_WRONG: "token不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式不正确",
	ERROR_USER_NO_RIGHT: "用户无权限",

	// 分类
	ERROR_CATENAME_USED: "分类名称已存在",

	ERROR_CATE_NOT_EXIST: "该分类不存在",
}

func GetErrMsg(code int) string {
        return codeMsg[code]
}

