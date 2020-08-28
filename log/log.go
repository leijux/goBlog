package log

import (
	"io"
	"log"
	"os"
	"path"
	"time"

	"goBlog/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//Logger 全局的log记录器
var Logger *logrus.Logger

func init() {
	logFilePath := config.GetString("log.logFilePath")
	logFileName := config.GetString("log.logFileName")

	fileName := path.Join(logFilePath, logFileName) // 日志文件

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println("err:", err)
	}

	Logger = logrus.New()                       // 实例化
	Logger.Out = io.MultiWriter(src, os.Stderr) // 设置输出true
	Logger.SetReportCaller(true)                //记录 行号 文件和函数
	
	Logger.SetLevel(logrus.DebugLevel)          // 设置日志级别 在什么级别之上

	logWriter, err := rotatelogs.New( // 设置 rotatelogs
		fileName+".%Y%m%d.log", // 分割后的文件名称

		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(fileName),

		rotatelogs.WithMaxAge(7*24*time.Hour), // 设置最大保存时间(7天)

		rotatelogs.WithRotationTime(24*time.Hour), // 设置日志切割时间间隔(1天)
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	//日志的等级
	// DebugLevel
	// InfoLevel
	// WarnLevel
	// ErrorLevel
	// FatalLevel
	// PanicLevel
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增钩子
	Logger.AddHook(lfHook)
}
