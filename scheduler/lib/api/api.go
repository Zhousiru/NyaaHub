package api

import (
	"github.com/Zhousiru/NyaaHub/scheduler/lib/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(config.SchedulerApiDebug)
}

func Start() {
	r := gin.Default()
	r.Use(authMiddleware)
	r.Use(cors.Default())

	r.POST("/addTask", addTask)
	r.POST("/updateTaskConfig", updateTaskConfig)
	r.GET("/removeTask", removeTask)
	r.GET("/listTask", listTask)
	r.GET("/getLog", getLog)

	go r.Run(config.SchedulerApiListen)
}
