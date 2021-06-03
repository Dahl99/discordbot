package discordbot

import (
	"encoding/json"
	"log"
	"net/http"
)

type ytPage struct {
	Items []items `json:"items"`
}

type items struct {
	Id id `json:"id"`
	Snippet snippet `json:"snippet"`
}

type id struct {
	VideoId string `json:"videoId"`
}

type snippet struct {
	Title string `json:"title"`
}

func ytSearch(url string) string {

	res, err := http.Get(url)
	if err != nil {
		log.Println(http.StatusServiceUnavailable)
		return ytSearchFailed
	}

	var page ytPage

	err = json.NewDecoder(res.Body).Decode(&page)
	if err != nil {
		log.Println(err)
		return decodingFailed
	}

	res.Body.Close()

	return page.Items[0].Id.VideoId + "|" + page.Items[0].Snippet.Title
}
