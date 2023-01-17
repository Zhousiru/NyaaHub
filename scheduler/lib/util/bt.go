package util

import (
	"net/url"
	"strings"
)

func GenMagnet(title, hash string) string {
	link := `magnet:?xt=urn:btih:{{hash}}&dn={{title}}&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce`
	link = strings.Replace(link, "{{hash}}", hash, 1)
	link = strings.Replace(link, "{{title}}", url.QueryEscape(title), 1)

	return link
}
