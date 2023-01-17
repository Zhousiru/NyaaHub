package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

type Task struct {
	Collection string
	Config     TaskConfig
	Downloaded int
	LastUpdate time.Time
}

type NewTask struct {
	Collection string
	Config     TaskConfig
}

type TaskConfig struct {
	Rss          string `json:"rss"`
	Cron         string `json:"cron"`
	CronTimeZone string `json:"cronTimeZone"`
	MaxDownload  int    `json:"maxDownload"`
	Timeout      int    `json:"timeout"`
}

func AddTask(task NewTask) error {
	taskConfigBytes, err := json.Marshal(task.Config)
	if err != nil {
		return err
	}
	taskConfig := string(taskConfigBytes)

	_, err = client.Exec(
		`INSERT INTO "task" ("collection", "config", "downloaded", "lastUpdate") VALUES (?, ?, ?, ?);`,
		task.Collection,
		taskConfig,
		0,
		time.Now().UTC(),
	)
	if err != nil {
		return err
	}

	return nil
}

func GetTask(collection string) (*Task, error) {
	row := client.QueryRow(
		`SELECT * FROM "task" WHERE "collection" = ?;`,
		collection,
	)

	return unmarshalRow(row)
}

func ExistTask(collection string) (bool, error) {
	_, err := GetTask(collection)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func RemoveTask(collection string) error {
	_, err := client.Exec(`DELETE FROM "task" WHERE "collection" = ?;`, collection)
	return err
}

func GetAllTask() ([]*Task, error) {
	rows, err := client.Query(`SELECT * FROM "task"`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return dumpTaskList(rows)
}

func GetAllTaskPagination(limit int, afterCollection string) ([]*Task, error) {
	rows, err := client.Query(`
	SELECT * FROM "task"
	WHERE "collection" > ?
	ORDER BY "collection"
	LIMIT ?
	`, afterCollection, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return dumpTaskList(rows)
}

func UpdateTaskConfig(collection string, new TaskConfig) error {
	taskConfigBytes, err := json.Marshal(new)
	if err != nil {
		return err
	}
	taskConfig := string(taskConfigBytes)

	_, err = client.Exec(`
	UPDATE "task"
	SET "config" = ?
	WHERE "collection" = ?;
	`, taskConfig, collection)

	return err
}

func IncreaseTaskDownloadedCount(collection string, n int) error {
	_, err := client.Exec(`
	UPDATE "task"
	SET "downloaded" = "downloaded" + ?
	WHERE "collection" = ?;
	`, n, collection)

	return err
}

func UpdateTaskLastUpdateDate(collection string) error {
	_, err := client.Exec(`
	UPDATE "task"
	SET "lastUpdate" = ?
	WHERE "collection" = ?;
	`, time.Now().UTC(), collection)

	return err
}

func dumpTaskList(rows *sql.Rows) ([]*Task, error) {
	var taskList []*Task

	for rows.Next() {
		task, err := unmarshalRow(rows)
		if err != nil {
			return nil, err
		}
		taskList = append(taskList, task)
	}

	return taskList, nil
}

type rowScanner interface {
	Scan(dest ...interface{}) error
}

func unmarshalRow(row rowScanner) (*Task, error) {
	ret := new(Task)
	var configString string

	err := row.Scan(&ret.Collection, &configString, &ret.Downloaded, &ret.LastUpdate)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(configString), &ret.Config)

	return ret, nil
}
