package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)
import (
	rlog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
)

func Logger() gin.HandlerFunc {
	filePath := "log/log.log"
	linkName := "lastestlog.log"
	scr,err := os.OpenFile(filePath,os.O_RDWR|os.O_CREATE,0755)
	if err != nil {
		fmt.Println("err:",err)
	}
	logger := logrus.New()

	logger.Out = scr

	// 日志分割
	logger.SetLevel(logrus.DebugLevel)
	logWriter,_ := rlog.New(
		filePath+"%Y%m%d.log",
		rlog.WithMaxAge(7*24*time.Hour),
		rlog.WithRotationTime(24*time.Hour),
		rlog.WithLinkName(linkName), // 建立软连接查看最新日志
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.WarnLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap,&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)

	// 日志格式
	return func(c *gin.Context){
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms",int(math.Ceil(float64(stopTime.Nanoseconds())/100000.0)))
		hostName,err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}

		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0{
			dataSize = 0
		}

		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":hostName,
			"status":statusCode,
			"SpendTime":spendTime,
			"Ip":clientIp,
			"Method": method,
			"Path":path,
			"DataSize":dataSize,
			"Agent":userAgent,
		})

		if len(c.Errors) > 0{
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		}else if statusCode >= 400 {
			entry.Warn()
		}else {
			entry.Info()
		}
	}
}
