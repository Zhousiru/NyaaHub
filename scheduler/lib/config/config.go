package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Fetcher struct {
	ApiUrl string
	Token  string
}

var (
	SchedulerApiToken  string
	SchedulerApiListen string
	SchedulerApiDebug  string
	FetcherConfig      *Fetcher
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

	// TODO: Support multiple fetcher
	fetcherConfigSplit := strings.Split(os.Getenv("FETCHER_CONFIG"), "|")
	FetcherConfig = &Fetcher{
		ApiUrl: fetcherConfigSplit[0],
		Token:  fetcherConfigSplit[1],
	}
}
