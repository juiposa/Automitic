package main

import (
	"github.com/juiposa/Automitic/internal/config"
	"github.com/juiposa/Automitic/internal/data"
	"github.com/juiposa/Automitic/internal/fflogs"
	"github.com/juiposa/Automitic/internal/logmanager"
	"github.com/juiposa/Automitic/internal/timeliner"
)

const ReportID = "nNcLBQqmKk9F6vMJ"

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	fflogsClient := fflogs.NewFFLogsClient(cfg)
	dataManager := data.InitDataManager(cfg)
	loader := logmanager.NewLogLoader(cfg, fflogsClient, dataManager)
	tlbuild := timeliner.NewTimeliner(cfg, loader, dataManager)
	timeline, err := tlbuild.BuildTimeline(ReportID, 2)
	if err != nil {
		panic(err)
	}
	err = dataManager.SaveResult(ReportID, 2, timeline)
	if err != nil {
		panic(err)
	}
}
