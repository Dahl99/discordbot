package music

import (
	"discordbot/src/config"
	"discordbot/src/consts"
	"discordbot/src/utils"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// func Music(cmd []string, v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) {

// 	if len(cmd) < 1 {
// 		return
// 	}

// 	switch(cmd[0]) {
// 	case "join":
// 		joinVoice(v, s, m)
// 	case "leave":
// 		leaveVoice(v, m)
// 	case "play":
// 		playMusic(cmd[1:], v, m)
// 	case "skip":
// 		skipMusic(v, m)
// 	case "stop":
// 		stopMusic(v, m)
// 	default:
// 		return
// 	}
// }

func JoinVoice(v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) {
	voiceChannelID := utils.SearchVoiceChannel(m.Author.ID)
	if voiceChannelID == "" {
		log.Println("Voice channel id not found")
		config.Dg.ChannelMessageSend(m.ChannelID, "[Music] You need to join a voice channel first!")
		return
	}

	if v != nil {
		log.Println("INFO: Voice instance already created")
	} else {
		guildID := utils.SearchGuild(m.ChannelID)
		mutex.Lock()
		v = new(VoiceInstance)
		VoiceInstances[guildID] = v
		v.GuildID = guildID
		v.session = s
		mutex.Unlock()
	}

	var err error
	v.voice, err = config.Dg.ChannelVoiceJoin(v.GuildID, voiceChannelID, false, true)
	if err != nil {
		v.Stop()
		log.Println("Error when joining voice channel")
		return
	}

	v.voice.Speaking(false)
	log.Println("New voice instance created")
	config.Dg.ChannelMessageSend(m.ChannelID, "[Music] Voice channel joined!")
}


func LeaveVoice(v *VoiceInstance, m *discordgo.MessageCreate) {
	log.Println("INFO:", m.Author.Username, "requested 'leave'")

	if v == nil {
		log.Println("INFO: The bot is not in a voice channel!")
		return
	}

	v.Stop()
	time.Sleep(200 * time.Millisecond)
	v.voice.Disconnect()
	log.Println("INFO: Voice channel left")
	mutex.Lock()
	delete(VoiceInstances, v.GuildID)
	mutex.Unlock()
	config.Dg.ChannelMessageSend(m.ChannelID, "[Music] Voice channel left!")
}


func PlayMusic(n []string, v *VoiceInstance, m *discordgo.MessageCreate) {
	if v == nil {
		log.Println("INFO: The bot is not in a voice channel")
		config.Dg.ChannelMessageSend(m.ChannelID, "[Music] I need to join a voice channel first!")
		return
	}

	voiceChannelID := utils.SearchVoiceChannel(m.Author.ID)
	if v.voice.ChannelID != voiceChannelID {
		config.Dg.ChannelMessageSend(m.ChannelID, "[Music] <@" + m.Author.ID + "> you need to join my voice channel first!")
		return
	}

	var videoId string
	var videoTitle string
	var err error

	// If a youtube url is sent as argument
	if strings.Contains(n[0], consts.YtVideoUrl) {
		url, err := url.Parse(n[0])
		if err != nil {
			log.Println("INFO: Unable to parse youtube url")
			config.Dg.ChannelMessageSend(m.ChannelID, "[Music] Oops, something wrong happened when parsing youtube url")
			return
		}

		query := url.Query()
		videoId = query.Get("v")

		videoTitle, err = ytFind(videoId)
		if err != nil {
			log.Println("INFO: unable to find title of song on youtube")
			config.Dg.ChannelMessageSend(m.ChannelID, "[Music] Oops, something went wrong when fetching title on YouTube")
			return
		}

	// If argument(s) is not a youtube url
	} else {
		name := utils.ReplaceSpace(n)
		videoId, videoTitle, err = ytSearch(name)
		if err != nil {
			log.Println("INFO: Unable to find song by searching youtube")
			config.Dg.ChannelMessageSend(m.ChannelID, "[Music] Oops, something wrong happened when searching YouTube")
			return
		}
	}

	song, err := execYtdl(videoId, videoTitle, v, m)

	if err != nil || song.data.ID == "" {
		log.Println("INFO: Youtube search: ", err)
		config.Dg.ChannelMessageSend(m.ChannelID, "[Music] Unable to find song")
		return
	}

	config.Dg.ChannelMessageSend(m.ChannelID, "[Music] " + song.data.User + " has added **" + song.data.Title + "** to the queue")

	go func() {
		SongSignal <- song
	}()
}

func SkipMusic(v *VoiceInstance, m *discordgo.MessageCreate) {
	log.Println("INFO:", m.Author.Username, "send 'skip'")
	if v == nil {
		log.Println("INFO: The bot is not in a voice channel")
		config.Dg.ChannelMessageSend(m.ChannelID, "[Music] I need to join a voice channel first!")
		return
	}
	if len(v.queue) == 0 {
		log.Println("INFO: The queue is empty.")
		config.Dg.ChannelMessageSend(m.ChannelID, "[Music] There is no song playing")
		return
	}
	if v.Skip() {
		config.Dg.ChannelMessageSend(m.ChannelID, "[Music] I'm paused, resume first")
	}
}


func StopMusic(v *VoiceInstance, m *discordgo.MessageCreate) {
	log.Println("INFO:", m.Author.Username, "requested stopping of music")

	if v == nil {
		log.Println("INFO: The bot is not in a voice channel")
		config.Dg.ChannelMessageSend(m.ChannelID, "[Music] I need to join a voice channel first!")
		return
	}
	voiceChannelID := utils.SearchVoiceChannel(m.Author.ID)
	if v.voice.ChannelID != voiceChannelID {
		config.Dg.ChannelMessageSend(m.ChannelID, "[Music] <@"+m.Author.ID+"> You need to join my voice channel to stop music!")
		return
	}

	v.Stop()
	log.Println("INFO: The bot stopped playing music")
	config.Dg.ChannelMessageSend(m.ChannelID, "[Music] I have now stopped playing music!")
}
