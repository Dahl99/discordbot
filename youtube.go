package discordbot

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	ytdl "github.com/kkdai/youtube/v2"
	"google.golang.org/api/option"
	ytapi "google.golang.org/api/youtube/v3"
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

func ytSearch(name string, v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) (song_struct PkgSong, err error) {

	// res, err := http.Get(url)
	// if err != nil {
	// 	log.Println(http.StatusServiceUnavailable)
	// 	return
	// }

	// var page ytPage

	// err = json.NewDecoder(res.Body).Decode(&page)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// ytapiclient := &http.Client {
	// 	Transport: &transport.APIKey{Key: conf.Ytkey},
	// }

	// service, err := ytapi.New(ytapiclient)
	// if err != nil {
	// 	return
	// }

	fmt.Println("name = ", name)

	ctx := context.Background()

	// client := ytapi.Client(ctx, conf.Ytkey)
	// client := ytapi.New(ctx, conf.Ytkey)
	service, err := ytapi.NewService(ctx, option.WithAPIKey(conf.Ytkey))
	if err != nil {
		return
	}

	args := []string{"contentDetails", "id", "snippet", "id,snippet"}

	fmt.Println("args = ", args[1:])

	call := service.Search.List(args[1:2]).Q(ytVideoUrl + name).MaxResults(1)
	res, err := call.Do()
	if err != nil {
		return
	}

	fmt.Println("res.items[0] = ", res.Items[0])

	var (
		audioId, audioTitle string
	)

	// fmt.Println("id = ", res.Items[0].Id.VideoId)
	// fmt.Println("title = ", res.Items[0].Snippet.Title)

	for _, item := range res.Items {
		audioId = item.Id.VideoId
		audioTitle = "item.Snippet.Title"
		// audioTitle = item.Snippet.Title
	}


	fmt.Println("id = ", audioId, " title = ", audioTitle)

	// audioId = page.Items[0].Id.VideoId
	// audioTitle = page.Items[0].Id.VideoId

	if audioId == "" {
		s.ChannelMessageSend(m.ChannelID, "[Music] Sorry, I'm unable to find the requested song")
		return
	}

	ytdlclient := ytdl.Client{}

	vid, err := ytdlclient.GetVideo(ytVideoUrl + audioId)
	if err != nil {
		return
	}

	fmt.Println("Formats length = ", len(vid.Formats))

	// format := vid.Formats[0]
	format := vid.Formats.
	videoURL, err := ytdlclient.GetStreamURL(vid, &format)
	if err != nil {
		return
	}

	videos := service.Videos.List(args[:0]).Id(vid.ID)
	videos.Do()

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
		vid.ID,
		audioTitle,
		videoURL,
	}

	song_struct.data = song
	song_struct.v = v

	// duration := res.Items[0].

	// res.Body.Close()

	return
}
