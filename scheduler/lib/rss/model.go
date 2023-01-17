package rss

type RssRes struct {
	Channel struct {
		Item []*RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Guid        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
	Seeders     string `xml:"seeders"`
	Leechers    string `xml:"leechers"`
	Downloads   string `xml:"downloads"`
	InfoHash    string `xml:"infoHash"`
	CategoryId  string `xml:"categoryId"`
	Category    string `xml:"category"`
	Size        string `xml:"size"`
	Comments    string `xml:"comments"`
	Trusted     string `xml:"trusted"`
	Remake      string `xml:"remake"`
	Description string `xml:"description"`
}
