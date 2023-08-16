package main

import (
	"github.com/korewayume/zlog"
	"go.uber.org/zap"
)

func main() {
	zlog.SetLevel(zap.ErrorLevel)
	zlog.Debug("1")
	zlog.Debug("2")
	zlog.SetLevel(zap.DebugLevel)
	zlog.Debug("3")
}
