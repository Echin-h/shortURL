package log

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	SugarLogger *zap.SugaredLogger
)

// 做一个闭包
func Init() gin.HandlerFunc {
	return func(c *gin.Context) {
		initLogger()
		c.Next()
	}
}

func initLogger() {
	Encoder := getEncoder()
	WriterSyncer := getWriterSyncer()
	core := zapcore.NewCore(Encoder, WriterSyncer, zapcore.DebugLevel)
	//  zap.AddCaller()可以实现记录函数信息
	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

// 编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 写入位置
func getWriterSyncer() zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/log.log",
		MaxSize:    1 << 30,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
	})
}
