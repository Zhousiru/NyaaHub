package bt

import (
	"path"
	"path/filepath"
	"time"

	"github.com/Zhousiru/NyaaHub/fetcher/lib/config"
)

func GetDownloadDir(collection string) string {
	suffix := " [" + time.Now().UTC().Format("2006-01-02 15:04:05.999") + "]"
	return path.Join(config.BTDownloadDir, collection+suffix)
}

func GetCollection(downloadDir string) (string, error) {
	dirname, err := filepath.Rel(
		config.BTDownloadDir,
		downloadDir,
	)
	if err != nil {
		return "", err
	}

	return dirname[:len(dirname)-26], nil
}
