package rss

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Zhousiru/NyaaHub/scheduler/lib/db"
)

func GetRssUpdate(task *db.Task) ([]*RssItem, error) {
	itemList, err := GetRss(task.Config.Rss)
	if err != nil {
		return nil, err
	}

	var newItemList []*RssItem

	for _, item := range itemList {
		pubDate, err := time.Parse(`Mon, 02 Jan 2006 15:04:05 -0700`, item.PubDate)
		if err != nil {
			return nil, err
		}
		if !pubDate.After(task.LastUpdate) {
			continue
		}
		newItemList = append(newItemList, item)
	}

	return newItemList, nil
}

func GetRss(rssUrl string) ([]*RssItem, error) {
	res, err := http.Get(rssUrl)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	rssRes := new(RssRes)

	err = xml.Unmarshal(bodyBytes, &rssRes)
	if err != nil {
		log.Fatal(err)
	}

	return rssRes.Channel.Item, nil
}
