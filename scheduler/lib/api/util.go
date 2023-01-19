package api

import (
	"errors"
	"strings"

	"github.com/Zhousiru/NyaaHub/scheduler/lib/db"
	"github.com/gin-gonic/gin"
)

func bindNewTask(newTask *db.NewTask, c *gin.Context) error {
	err := c.BindJSON(&newTask)
	if err != nil ||
		strings.TrimSpace(newTask.Collection) == "" {
		return errors.New("invalid params")
	}
	return nil
}
