package api

import (
	"net/http"
	"strings"

	"github.com/Zhousiru/NyaaHub/scheduler/lib/db"
	"github.com/gin-gonic/gin"
)

func getLog(c *gin.Context) {
	// TODO: Pagination?

	collection := c.Query("collection")
	if strings.TrimSpace(collection) == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"payload": nil,
				"msg":     "invalid params",
			},
		)
		return
	}

	logList, err := db.GetCollectionLog(collection)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"payload": nil,
				"msg":     err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"payload": logList,
			"msg":     "ok",
		},
	)
}
