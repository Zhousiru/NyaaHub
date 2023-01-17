package cronjob

import (
	"fmt"
	"sync"

	"github.com/Zhousiru/NyaaHub/scheduler/lib/db"
	"github.com/robfig/cron/v3"
)

var c = cron.New()
var idMap sync.Map

func UpdateCronJob(task *db.Task) error {
	prevID, ok := idMap.Load(task.Collection)
	if ok {
		c.Remove(prevID.(cron.EntryID))
		idMap.Delete(task.Collection)
	}

	cronStr := fmt.Sprintf("CRON_TZ=%s %s", task.Config.CronTimeZone, task.Config.Cron)
	id, err := c.AddFunc(cronStr, func() { CheckUpdate(task.Collection) })
	if err != nil {
		return err
	}

	idMap.Store(task.Collection, id)

	return nil
}

func RemoveCronJob(collection string) {
	id, ok := idMap.Load(collection)
	if ok {
		c.Remove(id.(cron.EntryID))
		idMap.Delete(collection)
	}
}
