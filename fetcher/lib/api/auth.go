package api

import (
	"net/http"

	"github.com/Zhousiru/NyaaHub/fetcher/lib/config"
	"github.com/gin-gonic/gin"
)

func authMiddleware(c *gin.Context) {
	token := c.Query("token")
	if token != config.FetcherApiToken {
		c.Status(http.StatusUnauthorized)
		c.Abort()
		return
	}

	c.Next()
}
