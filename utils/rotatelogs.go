package utils

import (
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
	"ziyue/global"
)

func GetWriterSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(global.Z_CONFIG.ZAP.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(global.Z_CONFIG.ZAP.LinkName),
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		zaprotatelogs.WithRotationTime(7*24*time.Hour),
	)
	if global.Z_CONFIG.ZAP.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
