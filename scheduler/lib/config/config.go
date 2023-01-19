package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	SchedulerApiToken  string
	SchedulerApiListen string
	SchedulerApiDebug  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	SchedulerApiToken = os.Getenv("SCHEDULER_API_TOKEN")
	SchedulerApiListen = os.Getenv("SCHEDULER_API_LISTEN")
	schedulerApiDebugBool, err := strconv.ParseBool(os.Getenv("SCHEDULER_API_DEBUG"))
	if err != nil {
		log.Fatalln(err)
	}
	if schedulerApiDebugBool {
		SchedulerApiDebug = "debug"
	} else {
		SchedulerApiDebug = "release"
	}
}
