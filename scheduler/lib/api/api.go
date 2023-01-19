package api

import (
	"github.com/Zhousiru/NyaaHub/scheduler/lib/config"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(config.SchedulerApiDebug)
}

func Start() {
	r := gin.Default()
	r.Use(authMiddleware)

	r.POST("/addTask", addTask)
	r.POST("/updateTaskConfig", updateTaskConfig)
	r.GET("/removeTask", removeTask)
	r.GET("/listTask", listTask)

	go r.Run(config.SchedulerApiListen)
}
