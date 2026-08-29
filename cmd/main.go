package main

import (
	"github.com/juiposa/Automitic/internal/config"
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/logmanager"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	fflogsClient := fflogs.NewFFLogsClient(cfg)
	loader := logmanager.NewLogLoader(cfg, fflogsClient)
	_, err = loader.LoadLogs("fpMxRBFWPgDXKCkA", 8)
	if err != nil {
		panic(err)
	}
}
