package bt

import (
	"context"
	"log"

	"github.com/Zhousiru/NyaaHub/fetcher/lib/config"
	"github.com/hekmon/transmissionrpc/v2"
)

var client *transmissionrpc.Client

func init() {
	var err error
	client, err = transmissionrpc.New(
		config.TransmissionRpcHost,
		config.TransmissionRpcUser,
		config.TransmissionRpcPasswd,
		&transmissionrpc.AdvancedConfig{
			HTTPS:  config.TransmissionRpcHttps,
			Port:   config.TransmissionRpcPort,
			RPCURI: config.TransmissionRpcUri,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
}

func AddTorrent(magnet, collection string) error {
	downloadDir := GetDownloadDir(collection)

	_, err := client.TorrentAdd(
		context.Background(),
		transmissionrpc.TorrentAddPayload{
			Filename:    &magnet,
			DownloadDir: &downloadDir,
		},
	)

	return err
}

func GetStatus() ([]transmissionrpc.Torrent, error) {
	torrents, err := client.TorrentGet(
		context.Background(),
		[]string{
			"id", "name", "eta", "magnetLink",
			"addedDate", "files", "percentDone", "downloadDir",
			"peersConnected", "sizeWhenDone", "status", "rateDownload",
			"rateUpload",
		},
		nil,
	)
	if err != nil {
		return nil, err
	}

	return torrents, nil
}

func RemoveTorrent(id int64) error {
	return client.TorrentRemove(
		context.Background(),
		transmissionrpc.TorrentRemovePayload{
			IDs:             []int64{id},
			DeleteLocalData: true,
		},
	)
}
