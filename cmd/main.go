package main

import (
	"github.com/juiposa/Automitic/internal/cache"
	"github.com/juiposa/Automitic/internal/config"
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/logmanager"
	"github.com/juiposa/Automitic/internal/timeliner"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	fflogsClient := fflogs.NewFFLogsClient(cfg)
	cacheManager := cache.NewCacheManager(cfg)
	loader := logmanager.NewLogLoader(cfg, fflogsClient, cacheManager)
	timeliner := timeliner.NewTimeliner(cfg, loader)
	_, err = timeliner.BuildTimeline("fpMxRBFWPgDXKCkA", 8)
	if err != nil {
		panic(err)
	}
}
