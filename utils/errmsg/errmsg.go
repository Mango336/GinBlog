/* 错误代码包 errmsg */
package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000...  用户模块错误
	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WORNG     = 1002
	ERROR_USERNAME_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST        = 1004
	ERROR_TOKEN_RUNTIME      = 1005
	ERROR_TOKEN_WRONG        = 1006
	ERROR_TOKEN_TYPE_WRONG   = 1007

	// code = 2000...  文章模块错误

	// code = 3000...  分类模块错误
	ERROR_CATEGORYNAME_USED = 3001
)

var errMp = map[int]string{
	SUCCESS:                  "OK",
	ERROR:                    "Error",
	ERROR_USERNAME_USED:      "用户名已存在...",
	ERROR_PASSWORD_WORNG:     "用户名或密码错误...",
	ERROR_USERNAME_NOT_EXIST: "用户名不存在...",
	ERROR_TOKEN_EXIST:        "TOKEN已存在...",
	ERROR_TOKEN_RUNTIME:      "TOKEN超时...",
	ERROR_TOKEN_WRONG:        "TOKEN错误...",
	ERROR_TOKEN_TYPE_WRONG:   "TOKEN格式错误...",

	ERROR_CATEGORYNAME_USED: "该分类已存在...",
}

func GetErrMsg(code int) string {
	return errMp[code]
}
