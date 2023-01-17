package db

import (
	"database/sql"
	"os"
	"path"
	"path/filepath"

	"github.com/Zhousiru/NyaaHub/scheduler/lib/util"
)

var dbPath string
var client *sql.DB

func init() {
	execFile, err := os.Executable()
	if err != nil {
		panic(err)
	}
	execDir := filepath.Dir(execFile)
	dbPath = path.Join(execDir, "data.db")

	dbNotExist := false
	if !util.PathExist(dbPath) {
		_, err := os.Create(dbPath)
		if err != nil {
			panic(err)
		}

		dbNotExist = true
	}

	client, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}

	if dbNotExist {
		CreateTable()
	}
}
