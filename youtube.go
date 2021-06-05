package discordbot

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"

	"github.com/bwmarrin/discordgo"
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

type videoResponse struct {
	Formats []struct {
		Url string `json:"url"`
	} `json:"formats"`
}

func ytSearch(name string, v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) (song_struct PkgSong, err error) {

	res, err := http.Get(youtubeEndpoint + conf.Ytkey + "&q=" + name)
	if err != nil {
		log.Println(http.StatusServiceUnavailable)
		return
	}

	var page ytPage

	err = json.NewDecoder(res.Body).Decode(&page)
	if err != nil {
		log.Println(err)
		return
	}

	videoID := page.Items[0].Id.VideoId
	videoTitle := page.Items[0].Snippet.Title

	cmd := exec.Command("youtube-dl", "--skip-download", "--print-json", "--flat-playlist", videoID)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		log.Println("ERROR: something wrong happened when running youtube-dl")
		return
	}

	var videoRes videoResponse
	err = json.NewDecoder(&out).Decode(&videoRes)

	guildID := searchGuild(m.ChannelID)
	member, _ := v.session.GuildMember(guildID, m.Author.ID)
	userName := ""

	if member.Nick == "" {
		userName = m.Author.Username
	} else {
		userName = member.Nick
	}

	song := Song {
		m.ChannelID,
		userName,
		m.Author.ID,
		videoID,
		videoTitle,
		videoRes.Formats[0].Url,
	}

	song_struct.data = song
	song_struct.v = v

	res.Body.Close()

	return
}
