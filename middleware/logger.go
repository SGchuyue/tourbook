package middleware

import (
	"fmt"
	sl "github.com/SGchuyue/logger/logger"
	"github.com/gin-gonic/gin"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	// InitLogger 初始化日志记录包
	sl.InitLogger("./test.log", 2, 3, 5, true)
	// 日志格式
	return func(c *gin.Context) {
		startTime := time.Now() // 开始时间
		c.Next()
		stopTime := time.Since(startTime)                                                           // 结束时间
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/100000.0))) // 执行时常
		hostName, err := os.Hostname()                                                              // 登陆主机名
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()    // HTTP响应状态码
		clientIp := c.ClientIP()           // 客地址ip
		userAgent := c.Request.UserAgent() // 客户端响应地址
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method   // 响应方式
		path := c.Request.RequestURI // 响应路径
		if len(c.Errors) == 0 {
			sl.Debug(spendTime, hostName, statusCode, clientIp, userAgent, method, path)
		}
		if len(c.Errors) > 0 {
			sl.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			sl.Warn(spendTime, hostName, statusCode, clientIp, userAgent, method, path)
		} else if statusCode >= 400 {
			sl.Info(spendTime, hostName, statusCode, clientIp, userAgent, method, path)
		}
	}
}
