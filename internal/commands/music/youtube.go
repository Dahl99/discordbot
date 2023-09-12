package music

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/api/option"

	"github.com/Dahl99/discord-bot/internal/config"
	"github.com/Dahl99/discord-bot/internal/discord"

	"github.com/bwmarrin/discordgo"
	ytdl "github.com/kkdai/youtube/v2"
	"google.golang.org/api/youtube/v3"
)

// youtubeFindEndpoint contains endpoint for finding more details about a video
const youtubeFindEndpoint string = "https://www.googleapis.com/youtube/v3/videos?part=snippet&key="

type Video struct {
	VideoID    string
	VideoTitle string
}

func SearchByName(name string) (*Video, error) {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(config.GetYoutubeApiKey()))

	call := youtubeService.Search.List([]string{"id", "snippet"}).Q(name).MaxResults(1)
	res, err := call.Do()
	if err != nil {
		slog.Warn("failed to search YouTube for video", "error", err)
		return nil, err
	}

	var (
		videoID, videoTitle string
	)

	for _, item := range res.Items {
		videoID = item.Id.VideoId
		videoTitle = item.Snippet.Title
	}

	if videoID == "" {
		slog.Info("video not found on YouTube")
		return nil, errors.New("video not found on YouTube")
	}

	return &Video{
		VideoID:    videoID,
		VideoTitle: videoTitle,
	}, nil
}

func findByVideoID(videoID string) (*Video, error) {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(config.GetYoutubeApiKey()))

	call := youtubeService.Videos.List([]string{"id", "snippet"}).Id(videoID).MaxResults(1)
	res, err := call.Do()
	if err != nil {
		slog.Warn("failed to find video on YouTube", "error", err)
		return nil, err
	}

	var videoTitle string

	for _, item := range res.Items {
		videoTitle = item.Snippet.Title
	}

	if videoID == "" {
		slog.Info("video not found on YouTube")
		return nil, errors.New("video not found on YouTube")
	}

	return &Video{
		VideoID:    videoID,
		VideoTitle: videoTitle,
	}, nil
}

func execYtdl(video *Video, v *VoiceInstance, m *discordgo.MessageCreate) (songStruct PkgSong, err error) {
	client := ytdl.Client{
		Debug:       false,
		HTTPClient:  nil,
		MaxRoutines: 0,
		ChunkSize:   0,
	}

	youTubeVideo, err := client.GetVideo("https://www.youtube.com/watch?v=" + video.VideoID)
	if err != nil {
		slog.Warn("failed to get video from YouTube", "error", err)
		return
	}

	youTubeVideo.Formats.Sort()
	formatURL := youTubeVideo.Formats[0].URL

	guildID := discord.SearchGuildByChannelID(m.ChannelID)
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
		video.VideoID,
		video.VideoTitle,
		formatURL,
	}

	// var song_struct PkgSong
	songStruct.data = song
	songStruct.v = v

	return songStruct, nil
}
