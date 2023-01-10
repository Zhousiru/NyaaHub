package api

import (
	"net/http"

	"github.com/Zhousiru/NyaaHub/fetcher/lib/bt"
	"github.com/gin-gonic/gin"
)

type resp struct {
	Payload any    `json:"payload"`
	Msg     string `json:"msg"`
}

func handleStatus(c *gin.Context) {
	torrents, err := bt.GetStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp{
			Payload: nil,
			Msg:     "failed to get bt status: " + err.Error(),
		})
		return
	}

	var payload []map[string]interface{}

	for _, t := range torrents {
		payload = append(payload, trimNil(t))
	}

	c.JSON(http.StatusOK, resp{
		Payload: payload,
		Msg:     "",
	})
}

func handleAdd(c *gin.Context) {
	magnet := c.Query("magnet")
	collection := c.Query("collection")
	if magnet == "" || collection == "" {
		c.JSON(http.StatusBadRequest, resp{
			Payload: nil,
			Msg:     "invalid params",
		})

		return
	}

	err := bt.AddTorrent(magnet, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp{
			Payload: nil,
			Msg:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp{
		Payload: nil,
		Msg:     "",
	})
}
