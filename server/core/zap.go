package core

import (
	"fmt"
	"os"

	"github.com/kesilent/react-light-blog/core/internal"
	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.RLB_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.RLB_CONFIG.Zap.Director)
		_ = os.Mkdir(global.RLB_CONFIG.Zap.Director, os.ModePerm)
	}
	levels := global.RLB_CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if global.RLB_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
