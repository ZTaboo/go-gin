package z

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Log *zap.SugaredLogger

func init() {
	InitLogger()

}

func InitLogger() {
	/*logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalln("初始化log组件错误：", err)
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	Log = logger.Sugar()
	*/
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	WriteSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(encoder, WriteSyncer, zapcore.DebugLevel)
	logger := zap.New(core)
	Log = logger.Sugar()
}
