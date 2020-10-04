// 登陆验证
package middleware

import (
	"tourbook/utils"
	"tourbook/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte(utils.JwtKey)
var code int

type MyClaims struct {
	Username string `json:"username"`
//	Password string `json: "password"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string,int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims {
		username,
//		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "tourbook",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaims) // hs256 非 ES256
	token,err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "",errmsg.ERROR
	}
	return token,errmsg.SUCCSE
}

// 验证Token
func CheckToken( token string) (*MyClaims,int) {
	setToken,_ := jwt.ParseWithClaims(token,&MyClaims{},func(token *jwt.Token)(interface{},error){
		return JwtKey,nil
	})
	if key,_ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key,errmsg.SUCCSE
	}else{
		return nil,errmsg.ERROR
	}
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		code = errmsg.SUCCSE
		if tokenHerder == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK,gin.H{
                        "code": code,
                        "message": errmsg.GetErrMsg(code),
                })

			c.Abort()
			return
		}

		checkToken := strings.SplitN(tokenHerder,"",2)
		if len(checkToken) != 2&& checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
                        "code": code,
                        "message": errmsg.GetErrMsg(code),
                })
			c.Abort()
		}
		key,Tcode := CheckToken(checkToken[1])
		if Tcode == errmsg.ERROR{
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK,gin.H{
                        "code": code,
                        "message": errmsg.GetErrMsg(code),
                })

			c.Abort()
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK,gin.H{
                        "code": code,
                        "message": errmsg.GetErrMsg(code),
                })
			c.Abort()
		}
		/*
		c.JSON(http.StatusOK,gin.H{
			"code": code,
			"message": errmsg.GetErrMsg(code),
		})*/
		c.Set("username",key.Username)
		c.Next()

	}
}
