package feed_test

import (
	"fmt"
	"github.com/nicr9/feed"
)

// Examples

func ExampleNew() {
	data := []byte(`<rss>
		<channel>
			<title>Test RSS feed.</title>
		</channel>
	</rss>`)
	rss, err := feed.New(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", rss)
}

func ExampleNewFromHttp() {
	rss, err := feed.NewFromHttp("http://www.training.dupont.de/playlists/rss.xml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", rss)
}

func ExampleNewFromFile() {
	rss, err := feed.NewFromFile("./podcastinit.xml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", rss)
}

func ExampleToFile() {
	rss, err := feed.NewFromFile("./podcastinit.xml")
	if err != nil {
		fmt.Println(err)
	}
	rss.ToFile("./podcastinit.xml.bak")
}
