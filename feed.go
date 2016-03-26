package feed

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type FeedRss struct {
	Channel FeedChannel `xml:"channel"`
}

func (f FeedRss) ToFile(path string) error {
	body, err := xml.Marshal(f)
	if err != nil {
		return err
	}

	ioutil.WriteFile(path, body, 0755)

	return nil
}

type FeedChannel struct {
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	PubDate     string     `xml:"pubDate"`
	Ttl         int        `xml:"ttl"`
	Items       []FeedItem `xml:"item"`
}

type FeedItem struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	PubDate     string        `xml:"pubDate"`
	Description string        `xml:"description"`
	Enclosure   FeedEnclosure `xml:"enclosure"`
}

type FeedEnclosure struct {
	Url    string `xml:"url,attr"`
	Length int    `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func New(data []byte) (FeedRss, error) {
	feed := FeedRss{}
	err := xml.Unmarshal(data, &feed)
	if err != nil {
		return feed, err
	}
	return feed, nil
}

func NewFromHttp(url string) (FeedRss, error) {
	// Read bytes from http request
	data, err := httpGet(url)
	if err != nil {
		return FeedRss{}, err
	}

	// Unmarshal to FeedRss{}
	feed, err := New(data)
	if err != nil {
		return feed, err
	}

	return feed, nil
}

func NewFromFile(path string) (FeedRss, error) {
	// Read bytes from file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return FeedRss{}, err
	}

	// Unmarshal to FeedRss{}
	feed, err := New(data)
	if err != nil {
		return feed, err
	}

	return feed, nil
}
