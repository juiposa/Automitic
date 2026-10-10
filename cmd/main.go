package main

import (
	"github.com/juiposa/Automitic/internal/cache"
	"github.com/juiposa/Automitic/internal/config"
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/logmanager"
	"github.com/juiposa/Automitic/internal/timeliner"
)

const ReportID = "fpMxRBFWPgDXKCkA"

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	fflogsClient := fflogs.NewFFLogsClient(cfg)
	cacheManager := cache.NewCacheManager(cfg)
	loader := logmanager.NewLogLoader(cfg, fflogsClient, cacheManager)
	tlbuild := timeliner.NewTimeliner(cfg, loader)
	timeline, err := tlbuild.BuildTimeline(ReportID, 8)
	if err != nil {
		panic(err)
	}
	err = cacheManager.SaveResult(ReportID, 8, timeline)
	if err != nil {
		panic(err)
	}
}
