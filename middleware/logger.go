package middleware

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotalogs "github.com/lestrrat-go/file-rotatelogs" // 日志分割包
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	logPath := "log/" // 日志文件存储路径
	linkName := "latest_log.log"
	scr, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755) // 0755是文件的权限
	if err != nil {
		fmt.Println("err:", err)
	}

	logger := logrus.New()
	logger.Out = scr
	logger.SetLevel(logrus.DebugLevel) // 日志权限
	logWriter, _ := rotalogs.New(
		logPath+"%Y%m%d.log",                    // 文件名 年月日命名
		rotalogs.WithMaxAge(7*24*time.Hour),     // 保留时间 7天
		rotalogs.WithRotationTime(24*time.Hour), // 文件保存间隔时间
		rotalogs.WithLinkName(linkName),         // 软连接 连接到最新的日志文件
	)
	writerMap := lfshook.WriterMap{ // 将以下级别的log信息写入文件logWriter
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.TraceLevel: logWriter,
	}
	hook := lfshook.NewHook(writerMap, &logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logger.AddHook(hook)

	return func(c *gin.Context) {
		startTime := time.Now() // 客户端运行开始时间
		c.Next()                // 继续先执行后续的中间件
		endTime := time.Since(startTime)
		useTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(endTime.Nanoseconds()/1000.0))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknow"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize <= 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"HostName": hostName,
			"Status":   statusCode,
			"RunTime":  useTime,
			"Ip":       clientIp,
			"Method":   method,
			"DataSize": dataSize,
			"Path":     path,
			"Agent":    userAgent,
		})
		if len(c.Errors) > 0 { // 报错
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 { // warning code
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
