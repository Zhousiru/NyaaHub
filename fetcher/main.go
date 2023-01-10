package main

import (
	"log"
	"sync"

	"github.com/Zhousiru/NyaaHub/fetcher/lib/api"
	"github.com/Zhousiru/NyaaHub/fetcher/lib/bt"
	"github.com/Zhousiru/NyaaHub/fetcher/lib/config"
)

func main() {
	wg := new(sync.WaitGroup)

	api.StartApi(wg)
	log.Println("main: NyaaHub Fetcher API is listening on", config.FetcherApiListen)

	bt.StartWatch(wg)
	log.Println("main: NyaaHub Fetcher Watcher is running")

	wg.Wait()
}
