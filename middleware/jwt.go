package middleware

import (
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte(utils.JwtKey) // my signing key

type MyCustomClaims struct {
	Username string `json: username`
	// Password string `json: password`
	jwt.RegisteredClaims
}

// 生成token -- Create the claims
func SetToken(username string) (int, string) {
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 设置过期时间 24h
			Issuer:    "Mango",                                            // 签发人
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(JwtKey) // 签名字符串JWT
	if err != nil {
		return errmsg.ERROR, ""
	}
	return errmsg.SUCCESS, tokenStr
}

// 验证token -- parse with claims
func CheckToken(tokenStr string) (int, *MyCustomClaims) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return errmsg.ERROR, nil
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return errmsg.SUCCESS, claims // The token is true. token验证正确
	} else {
		return errmsg.ERROR, nil
	}
}

// jwt中间件
func JwtTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization") // 查看头部中是否有Authorization
		var code int = errmsg.SUCCESS
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort() // break 不再继续调用后续中间件
			return
		}
		token := strings.SplitN(tokenHeader, " ", 2)
		// token 格式不正确
		if len(token) != 2 && token[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkCode, claims := CheckToken(token[1])
		// token 验证错误
		if checkCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		// token超时
		if time.Now().Unix() > claims.ExpiresAt.Time.Unix() {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
		c.Set("username", claims.Username)
		c.Next() // 先继续调用后续中间件函数
	}
}
