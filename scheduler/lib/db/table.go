package db

import (
	_ "github.com/mattn/go-sqlite3"
)

func CreateTable() error {
	_, err := client.Exec(`
	CREATE TABLE "task" (
		"collection" TEXT NOT NULL,
		"config" TEXT NOT NULL,
		"downloaded" INTEGER NOT NULL,
		"lastUpdate" DATETIME NOT NULL,
		PRIMARY KEY ("collection")
	);
	`)
	if err != nil {
		return err
	}

	_, err = client.Exec(`
	CREATE TABLE "log" (
		"collection" TEXT NOT NULL,
		"time" DATETIME NOT NULL,
		"type" VARCHAR(50) NOT NULL,
		"msg" TEXT NOT NULL
	);
	`)
	if err != nil {
		return err
	}

	return nil
}
