package watcher

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Zhousiru/NyaaHub/fetcher/lib/bt"
	"github.com/Zhousiru/NyaaHub/fetcher/lib/config"
	"github.com/Zhousiru/NyaaHub/fetcher/lib/storage"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hekmon/transmissionrpc/v2"
)

func StartWatcher(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		watch()
	}()
}

func watch() {
	uploading := mapset.NewSet[int64]()

	for {
		time.Sleep(5 * time.Second)

		torrents, err := bt.GetStatus()
		if err != nil {
			log.Printf("watcher: failed to get bt status: %s", err)
			continue
		}

		for _, t := range torrents {
			if *t.Status != transmissionrpc.TorrentStatusSeedWait &&
				*t.Status != transmissionrpc.TorrentStatusSeed {
				continue
			}

			if uploading.Contains(*t.ID) {
				continue
			}

			relpath, err := filepath.Rel(
				config.BTDownloadDir,
				*t.DownloadDir,
			)
			if err != nil || strings.HasPrefix(relpath, "../") {
				continue
			}

			go func(torrent transmissionrpc.Torrent) {
				uploading.Add(*torrent.ID)
				defer uploading.Remove(*torrent.ID)

				log.Println("watcher: download complete:", *torrent.Name)

				downloadDir := *torrent.DownloadDir

				collection, err := bt.GetCollection(downloadDir)
				if err != nil {
					log.Printf("watcher: failed to get colletion by relpath: %s", err)
					return
				}

				log.Println("watcher: uploading:", *torrent.Name)

				err = storage.UploadDir(downloadDir, path.Join(config.RcloneRootDir, collection))
				if err != nil {
					log.Printf("watcher: failed to upload to storage: %s", err)
					return
				}

				log.Println("watcher: uploaded:", *torrent.Name)

				err = bt.RemoveTorrent(*torrent.ID)
				if err != nil {
					log.Printf("watcher: failed to reomove torrent: %s", *torrent.Name)
					return
				}

				dir, err := os.ReadDir(downloadDir)
				if err != nil {
					log.Printf("watcher: failed to read torrent dir: %s", downloadDir)
					return
				}

				if len(dir) == 0 {
					err = os.RemoveAll(downloadDir)
					if err != nil {
						log.Printf("watcher: failed to reomove torrent dir: %s", downloadDir)
						return
					}
				}
			}(t)
		}
	}
}
