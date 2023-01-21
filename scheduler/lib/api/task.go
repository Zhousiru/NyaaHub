package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Zhousiru/NyaaHub/scheduler/lib/cronjob"
	"github.com/Zhousiru/NyaaHub/scheduler/lib/db"
	"github.com/gin-gonic/gin"
)

func addTask(c *gin.Context) {
	newTask := db.NewTask{}
	err := bindNewTask(&newTask, c)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"payload": nil,
				"msg":     err.Error(),
			},
		)
		return
	}

	downloadOnly, _ := strconv.ParseBool(c.Query("downloadOnly"))

	if downloadOnly {
		cronjob.DownloadPrev(newTask.Collection, newTask.Config.Rss)
		c.JSON(http.StatusOK, gin.H{
			"payload": nil,
			"msg":     "ok",
		})
		return
	}

	downloadPrev, _ := strconv.ParseBool(c.Query("prev"))

	exist, err := db.ExistTask(newTask.Collection)
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

	if exist {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"payload": nil,
				"msg":     "specified `collection` already exists",
			},
		)
		return
	}

	err = db.AddTask(newTask)
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

	err = cronjob.UpdateCronJob(newTask.Collection, newTask.Config)
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

	c.JSON(http.StatusOK, gin.H{
		"payload": nil,
		"msg":     "ok",
	})

	if downloadPrev {
		cronjob.DownloadPrev(newTask.Collection, newTask.Config.Rss)
	}
}

func removeTask(c *gin.Context) {
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

	err := db.RemoveTask(collection)
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

	cronjob.RemoveCronJob(collection)

	c.JSON(http.StatusOK, gin.H{
		"payload": nil,
		"msg":     "ok",
	})
}

func updateTaskConfig(c *gin.Context) {
	task := db.NewTask{}
	err := bindNewTask(&task, c)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"payload": nil,
				"msg":     err.Error(),
			},
		)
		return
	}

	exist, err := db.ExistTask(task.Collection)
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

	if !exist {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"payload": nil,
				"msg":     "specified `collection` does not exist",
			},
		)
		return
	}

	err = db.UpdateTaskConfig(task.Collection, task.Config)
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

	err = cronjob.UpdateCronJob(task.Collection, task.Config)
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

	c.JSON(http.StatusOK, gin.H{
		"payload": nil,
		"msg":     "ok",
	})
}

func listTask(c *gin.Context) {
	start := c.Query("start")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"payload": nil,
				"msg":     "invalid params",
			},
		)
		return
	}

	taskList, err := db.GetAllTaskPagination(limit+1, start)
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

	hasNext := true
	if len(taskList) < limit+1 {
		hasNext = false
	}

	respTaskList := taskList
	if hasNext {
		respTaskList = taskList[:len(taskList)-1]
	}

	c.JSON(http.StatusOK,
		gin.H{
			"payload": gin.H{
				"list":    respTaskList,
				"hasNext": hasNext,
			},
			"msg": "ok",
		},
	)
}
