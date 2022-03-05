package music

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os/exec"

	"discordbot/src/config"
	"discordbot/src/utils"

	"github.com/bwmarrin/discordgo"
)

//	youtubeSearchEndpoint contains YouTube endpoint for searching after a video
const youtubeSearchEndpoint string = "https://www.googleapis.com/youtube/v3/search?part=snippet&type=video&key="

//	youtubeFindEndpoint contains endpoint for finding more details about a video
const youtubeFindEndpoint string = "https://www.googleapis.com/youtube/v3/videos?part=snippet&key="

// Structs for doing a Youtube search
type ytPageSearch struct {
	Items []itemsSearch `json:"items"`
}

type itemsSearch struct {
	Id      id      `json:"id"`
	Snippet snippet `json:"snippet"`
}

type id struct {
	VideoId string `json:"videoId"`
}

type snippet struct {
	Title string `json:"title"`
}

type videoResponse struct {
	Formats []struct {
		Url string `json:"url"`
	} `json:"formats"`
}

// Structs for finding a video on youtube
type ytPageFind struct {
	Items []itemsFind `json:"items"`
}

type itemsFind struct {
	Snippet snippet `json:"snippet"`
}

func ytSearch(name string) (string, string, error) {

	res, err := http.Get(youtubeSearchEndpoint + config.GetYoutubeKey() + "&q=" + name)
	if err != nil {
		log.Println(http.StatusServiceUnavailable)
		return "", "", err
	}

	var page ytPageSearch

	err = json.NewDecoder(res.Body).Decode(&page)
	if err != nil {
		log.Println(err)
		return "", "", err
	}

	res.Body.Close()

	if len(page.Items) < 1 {
		log.Println("INFO: empty youtube search result")
		err = errors.New("empty youtube search result")
		return "", "", err
	}
	videoId := page.Items[0].Id.VideoId
	videoTitle := page.Items[0].Snippet.Title

	return videoId, videoTitle, nil
}

func ytFind(videoId string) (string, error) {
	res, err := http.Get(youtubeFindEndpoint + config.GetYoutubeKey() + "&id=" + videoId)
	if err != nil {
		log.Println(http.StatusServiceUnavailable)
		return "", err
	}

	var page ytPageFind

	err = json.NewDecoder(res.Body).Decode(&page)
	if err != nil {
		log.Println(err)
		return "", err
	}

	res.Body.Close()

	if len(page.Items) < 1 {
		log.Println("INFO: empty youtube search result")
		err = errors.New("empty youtube search result")
		return "", err
	}

	videoTitle := page.Items[0].Snippet.Title

	return videoTitle, nil
}

func execYtdl(videoId string, videoTitle string, v *VoiceInstance, m *discordgo.MessageCreate) (songStruct PkgSong, err error) {

	cmd := exec.Command("youtube-dl", "--skip-download", "--print-json", "--flat-playlist", videoId)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		log.Println("ERROR: something wrong happened when running youtube-dl")
		return
	}

	var videoRes videoResponse
	err = json.NewDecoder(&out).Decode(&videoRes)
	if err != nil {
		log.Println("ERROR: error occurred when decoding video response")
		return
	}

	guildID := utils.SearchGuild(m.ChannelID)
	member, _ := v.session.GuildMember(guildID, m.Author.ID)
	userName := ""

	if member.Nick == "" {
		userName = m.Author.Username
	} else {
		userName = member.Nick
	}

	song := Song{
		m.ChannelID,
		userName,
		m.Author.ID,
		videoId,
		videoTitle,
		videoRes.Formats[2].Url,
	}

	// var song_struct PkgSong
	songStruct.data = song
	songStruct.v = v

	return songStruct, nil
}
