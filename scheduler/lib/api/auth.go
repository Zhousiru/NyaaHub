package api

import (
	"net/http"

	"github.com/Zhousiru/NyaaHub/scheduler/lib/config"
	"github.com/gin-gonic/gin"
)

func authMiddleware(c *gin.Context) {
	token := c.Query("token")
	if token != config.SchedulerApiToken {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	c.Next()
}
