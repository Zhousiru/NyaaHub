package bt

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/Zhousiru/NyaaHub/fetcher/lib/config"
	"github.com/Zhousiru/NyaaHub/fetcher/lib/storage"
	"github.com/hekmon/transmissionrpc/v2"
)

func StartWatch(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		watch()
	}()
}

func watch() {
	for {
		time.Sleep(5 * time.Second)

		torrents, err := GetStatus()
		if err != nil {
			log.Printf("watch: failed to get bt status: %s", err)
			continue
		}

		for _, t := range torrents {
			if *t.Status != transmissionrpc.TorrentStatusSeedWait &&
				*t.Status != transmissionrpc.TorrentStatusSeed {
				continue
			}

			log.Println("watch: download complete:", *t.Name)

			downloadDir := *t.DownloadDir

			successFlag := true

			for _, file := range t.Files {
				log.Println("watch: upload:", file.Name)

				downloadFilepath := path.Join(downloadDir, file.Name)

				collection, err := filepath.Rel(
					config.BTDownloadDir,
					downloadDir,
				)
				if err != nil {
					log.Printf("watch: failed to get colletion by relpath: %s", err)
					successFlag = false
					continue
				}

				err = storage.Upload(downloadFilepath, path.Join(config.RcloneRootDir, collection))
				if err != nil {
					log.Printf("watch: failed to upload to storage: %s", err)
					successFlag = false
					continue
				}
			}

			if successFlag {
				err = RemoveTorrent(*t.ID)
				if err != nil {
					log.Printf("watch: failed to remove torrent: %s", err)
					continue
				}

				err = os.Remove(downloadDir)
				if err != nil {
					log.Printf("watch: failed to remove torrent download dir: %s", err)
					continue
				}
			}
		}
	}
}
