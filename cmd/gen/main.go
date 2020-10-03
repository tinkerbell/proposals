package main

import (
	"github.com/tinkerbell/proposals/cmd/gen/cmd"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	cmd.Execute(logger)
}
