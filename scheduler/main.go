package main

import (
	"fmt"
	"os"

	"github.com/Zhousiru/NyaaHub/scheduler/lib/api"
	"github.com/Zhousiru/NyaaHub/scheduler/lib/config"
	"github.com/Zhousiru/NyaaHub/scheduler/lib/cronjob"
)

func main() {
	fmt.Println("[Cron Job] Load cron jobs from DB")
	err := cronjob.LoadAndStart()
	if err != nil {
		fmt.Println("[Cron Job] Failed:", err)
		os.Exit(1)
	}

	fmt.Println("[API] Listen on", config.SchedulerApiListen)
	api.Start()

	select {}
}
