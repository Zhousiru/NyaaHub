package cronjob

import (
	"fmt"
	"sync"

	"github.com/Zhousiru/NyaaHub/scheduler/lib/db"
	"github.com/robfig/cron/v3"
)

var c = cron.New()
var idMap sync.Map

func UpdateCronJob(collection string, config db.TaskConfig) error {
	prevID, ok := idMap.Load(collection)
	if ok {
		c.Remove(prevID.(cron.EntryID))
		idMap.Delete(collection)
	}

	cronStr := fmt.Sprintf("CRON_TZ=%s %s", config.CronTimeZone, config.Cron)
	id, err := c.AddFunc(cronStr, func() { checkUpdate(collection) })
	if err != nil {
		return err
	}

	idMap.Store(collection, id)

	return nil
}

func RemoveCronJob(collection string) {
	id, ok := idMap.Load(collection)
	if ok {
		c.Remove(id.(cron.EntryID))
		idMap.Delete(collection)
	}
}

func LoadAndStart() error {
	taskList, err := db.GetAllTask()
	if err != nil {
		return err
	}

	for _, task := range taskList {
		err := UpdateCronJob(task.Collection, task.Config)
		if err != nil {
			return fmt.Errorf("load %s: %w", task.Collection, err)
		}
	}

	c.Start()
	return nil
}
