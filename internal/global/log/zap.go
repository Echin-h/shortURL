package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	SugarLogger *zap.SugaredLogger
)

// 做一个闭包
func Init() {
	//return func(c *gin.Context) {
	//	initLogger()
	//	c.Next()
	//}
	initLogger()
}

func initLogger() {
	Encoder := getEncoder()
	// 输出地方为控制台怎么写
	WriterSyncer := zapcore.Lock(os.Stdout)
	fmt.Println("log os.Stdout success")
	core := zapcore.NewCore(Encoder, WriterSyncer, zapcore.DebugLevel)
	//  zap.AddCaller()可以实现记录函数信息
	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
	fmt.Println("log init success")
}

// 编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 写入位置
//func getWriterSyncer() zapcore.WriteSyncer {
//	return zapcore.AddSync(&lumberjack.Logger{
//		Filename:   "logs/log.log",
//		MaxSize:    1 << 30,
//		MaxBackups: 3,
//		MaxAge:     7,
//		Compress:   true,
//	})
//}
