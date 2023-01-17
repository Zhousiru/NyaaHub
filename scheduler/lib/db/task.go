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
	Rss          string        `json:"rss"`
	Cron         string        `json:"cron"`
	CronTimeZone string        `json:"cronTimeZone"`
	MaxDownload  int           `json:"maxDownload"`
	Timeout      time.Duration `json:"timeout"`
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
