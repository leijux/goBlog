package log

import (
	"goBlog/config"
	"io"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {

	logFilePath := config.GetString("log.logFilePath")
	logFileName := config.GetString("log.logFileName")

	fileName := path.Join(logFilePath, logFileName) // 日志文件
	logWriter, _ := rotatelogs.New(                 // 设置 rotatelogs
		fileName+".%Y%m%d.log", // 分割后的文件名称

		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(fileName),

		rotatelogs.WithMaxAge(7*24*time.Hour), // 设置最大保存时间(7天)

		rotatelogs.WithRotationTime(24*time.Hour), // 设置日志切割时间间隔(1天)
	)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(io.MultiWriter(logWriter, os.Stderr)),
		zap.InfoLevel,
	)

	logger = zap.New(core, zap.AddCaller(),zap.AddCallerSkip(1))
	defer logger.Sync()

	sugar = logger.Sugar()

}

type log struct {
	sugar *zap.SugaredLogger
}

func NewLog() log {
	return log{
		sugar: sugar,
	}
}

func (l log) Printf(s string, i ...interface{}) {
	l.sugar.Info(s, i)
}

func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

func Debugln(args ...interface{}) {
	sugar.Debug(args...)
}

func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}
func Infoln(args ...interface{}) {
	sugar.Info(args...)
}
func Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}
func Warn(msg string, fields ...zapcore.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	logger.Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func Errorln(args ...interface{}) {
	sugar.Error(args...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	logger.Fatal(msg, fields...)
}

func Fatalf(template string, args ...interface{}) {
	sugar.Fatalf(template, args...)
}

func Fatalln(args ...interface{}) {
	sugar.Fatal(args...)
}
