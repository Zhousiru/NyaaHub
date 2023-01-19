package cronjob

import (
	"strconv"
	"time"

	"github.com/Zhousiru/NyaaHub/scheduler/lib/db"
	"github.com/Zhousiru/NyaaHub/scheduler/lib/fetcher"
	"github.com/Zhousiru/NyaaHub/scheduler/lib/logger"
	"github.com/Zhousiru/NyaaHub/scheduler/lib/rss"
	"github.com/Zhousiru/NyaaHub/scheduler/lib/util"
)

func checkUpdate(collection string) {
	logger.Info(collection, "started by cron job")

	// Check lifespan
	logger.Info(collection, "check lifespan")
	task, err := db.GetTask(collection)
	if err != nil {
		logger.Error(collection, err.Error())
		return
	}

	if task.Downloaded >= task.Config.MaxDownload {
		logger.Info(collection, "max download exceeded. remove task")
		err := db.RemoveTask(collection)
		if err != nil {
			logger.Error(collection, err.Error())
			return
		}
		RemoveCronJob(collection)
		return
	}

	dur := time.Now().UTC().Sub(task.LastUpdate)

	if dur > time.Duration(task.Config.Timeout)*24*time.Hour {
		logger.Info(collection, "timeout exceeded. remove task")
		err := db.RemoveTask(collection)
		if err != nil {
			logger.Error(collection, err.Error())
			return
		}
		RemoveCronJob(collection)
		return
	}

	// Get RSS updates
	logger.Info(collection, "get rss updates")
	rssItemList, err := rss.GetRssUpdate(task)
	if err != nil {
		logger.Error(collection, err.Error())
		return
	}

	listLen := len(rssItemList)
	if listLen == 0 {
		logger.Info(collection, "no updates. exit")
		return
	}
	logger.Info(collection, "found "+strconv.Itoa(listLen)+" update(s)")

	// Download RSS updates
	for _, item := range rssItemList {
		logger.Info(collection, "download "+item.Title)
		fetcher.DownloadMagnet(collection, util.GenMagnet(item.Title, item.InfoHash))
	}
	db.IncreaseTaskDownloadedCount(collection, listLen)
	db.UpdateTaskLastUpdateDate(collection)

	logger.Info(collection, "all done. exit")
}
