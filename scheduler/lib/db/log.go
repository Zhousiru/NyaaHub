package db

import (
	"errors"
	"time"
)

type Log struct {
	Collection string    `json:"collection"`
	Time       time.Time `json:"time"`
	Type       string    `json:"type"`
	Msg        string    `json:"msg"`
}

const (
	LogTypeInfo = "info"
	LogTypeErr  = "error"
)

func AddLog(collection, logType, msg string) {
	client.Exec(`INSERT INTO "log" ("collection", "time", "type", "msg") VALUES (?, ?, ?, ?);`,
		collection, time.Now().UTC(), logType, msg,
	)
}

func GetCollectionLog(collection string) ([]*Log, error) {
	rows, err := client.Query(
		`SELECT * FROM "log" WHERE "collection" = ?;`,
		collection,
	)
	if err != nil {
		return nil, err
	}

	var logList []*Log

	for rows.Next() {
		log := new(Log)
		err := rows.Scan(&log.Collection, &log.Time, &log.Type, &log.Msg)
		if err != nil {
			return nil, err
		}
		logList = append(logList, log)
	}

	return logList, nil
}

func RemoveCollectionLog(collection string) error {
	// TODO
	return errors.New("not implemented")
}

func RemoveAllLog(collection string) error {
	// TODO
	return errors.New("not implemented")
}
