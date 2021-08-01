package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
	"ziyue/global"
	"ziyue/utils"
)

var level zapcore.Level

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.Z_CONFIG.ZAP.Director); !ok {
		fmt.Printf("Create %v directory\n", global.Z_CONFIG.ZAP.Director)
		_ = os.Mkdir(global.Z_CONFIG.ZAP.Director, os.ModePerm)
	}

	switch global.Z_CONFIG.ZAP.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if global.Z_CONFIG.ZAP.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.Z_CONFIG.ZAP.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.Z_CONFIG.ZAP.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.Z_CONFIG.ZAP.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色:
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.Z_CONFIG.ZAP.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.Z_CONFIG.ZAP.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

func getEncoder() zapcore.Encoder {
	if global.Z_CONFIG.ZAP.Format == "json" {
		zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderCore() (core zapcore.Core) {
	writer, err := utils.GetWriterSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err: %v", err.Error())
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.Z_CONFIG.ZAP.Prefix) + "2008/08/08-08:08:08.000")
}
