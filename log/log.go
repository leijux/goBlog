package log

import (
	"task-system/config"
	"fmt"

	"io"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {

	logFilePath := config.Cfg.Log.LOGFILEPATH
	logFileName := config.Cfg.Log.LOGFILENAME 

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	Logger = logrus.New()

	// 设置输出true
	Logger.Out = io.MultiWriter(src, os.Stderr)

	//记录 行号 文件和函数
	Logger.SetReportCaller(true)
	// 设置日志级别 在什么级别之上
	Logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
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
