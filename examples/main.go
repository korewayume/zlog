package main

import (
	"github.com/korewayume/zlog"
	"go.uber.org/zap"
)

func main() {
	zlog.SetLevel(zap.ErrorLevel)
	zlog.Debug("debug")
	zlog.SetLevel(zap.DebugLevel)
	zlog.Debug("debug")
	zlog.Error("error")
}
