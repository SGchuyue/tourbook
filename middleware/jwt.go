// 登陆验证
package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"tourbook/utils"
	"tourbook/utils/errmsg"
)

var JwtKey = []byte(utils.JwtKey)
var code int

type MyClaims struct {
	Username string `json:"username"`
	//	Password string `json: "password"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	ExpireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		username,
		//		password,
		jwt.StandardClaims{
			ExpiresAt: ExpireTime.Unix(),
			Issuer:    "tourbook",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims) // hs256 非 ES256
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCSE
}

// 验证Token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	key, _ := setToken.Claims.(*MyClaims)
	if setToken.Valid {
		return key, errmsg.SUCCSE
	} else {
		return nil, errmsg.ERROR
	}
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code = errmsg.SUCCSE
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, "", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
		}
		key, tcode := CheckToken(checkToken[1])
		if tcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
		}
		c.Set("username", key.Username)
		c.Next()

	}
}
