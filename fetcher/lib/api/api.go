package api

import (
	"log"
	"sync"

	"github.com/Zhousiru/NyaaHub/fetcher/lib/config"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(config.FetcherApiDebug)
}

func StartApi(wg *sync.WaitGroup) {
	r := gin.Default()
	r.Use(authMiddleware)

	r.GET("/status", handleStatus)
	r.GET("/add", handleAdd)

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := r.Run(config.FetcherApiListen)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
