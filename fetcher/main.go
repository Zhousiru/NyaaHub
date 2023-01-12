package main

import (
	"log"
	"sync"

	"github.com/Zhousiru/NyaaHub/fetcher/lib/api"
	"github.com/Zhousiru/NyaaHub/fetcher/lib/config"
	"github.com/Zhousiru/NyaaHub/fetcher/lib/watcher"
)

func main() {
	wg := new(sync.WaitGroup)

	api.StartApi(wg)
	log.Println("main: NyaaHub Fetcher API is listening on", config.FetcherApiListen)

	watcher.StartWatcher(wg)
	log.Println("main: NyaaHub Fetcher Watcher is running")

	wg.Wait()
}
