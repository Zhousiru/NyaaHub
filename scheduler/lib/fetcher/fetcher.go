package fetcher

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Zhousiru/NyaaHub/scheduler/lib/config"
)

func DownloadMagnet(collection, magnet string) error {
	req, err := http.NewRequest(
		http.MethodGet,
		config.FetcherConfig.ApiUrl+"/add",
		nil,
	)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("token", config.FetcherConfig.Token)
	q.Add("magnet", magnet)
	q.Add("collection", collection)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("failed to add download task for `" + collection + "`: fetcher response with status code " + strconv.Itoa(res.StatusCode))
	}

	return nil
}
