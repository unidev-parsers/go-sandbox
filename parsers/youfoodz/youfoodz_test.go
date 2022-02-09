package youfoodz

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/unidev-platform/golang-core/xstring"
)

func TestYoufoodzMainPageUrls(t *testing.T) {
	url := "https://youfoodz.com/collections/all"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	body := string(html)
	parts := xstring.Between(body, "collectionsSort['mains'].push('", "')")
	for _, item := range parts {
		log.Println(item)
	}
}
