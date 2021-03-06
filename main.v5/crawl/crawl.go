package crawl

import (
	"fmt"
	"strings"

	"./crawlphoto"
	"./crawlurl"
)

// access shops category of website walmart.com
func Crawl(urlSlice []string, urlMap map[string]bool) {
	fmt.Println(len(urlSlice))
	for _, url := range urlSlice {
		check := strings.Index(url, "photos3") != -1
		if check {
			crawlphoto.CrawlPhoto(url)
		} else {
			urlMap = crawlurl.CrawlUrl(url, urlMap)
		}
	}
}
