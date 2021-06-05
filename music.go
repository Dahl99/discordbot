package discordbot

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func music(cmd []string, v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) {

	if len(cmd) < 1 {
		return
	}

	switch(cmd[0]) {
	case "join":
		joinVoice(v, s, m)
	case "leave":
		leaveVoice(v, s, m)
	case "play":
		playMusic(cmd[1:], v, s, m)
	case "skip":
		skipMusic(s, m)
	default:
		return
	}
}

func joinVoice(v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) {
	voiceChannelID := searchVoiceChannel(m.Author.ID)
	if voiceChannelID == "" {
		log.Println("Voice channel id not found")
		s.ChannelMessageSend(m.ChannelID, "[Music] You need to join a voice channel first!")
		return
	}

	if v != nil {
		log.Println("INFO: Voice instance already created")
	} else {
		guildID := searchGuild(m.ChannelID)
		mutex.Lock()
		v = new(VoiceInstance)
		voiceInstances[guildID] = v
		v.guildID = guildID
		v.session = s
		mutex.Unlock()
	}

	var err error
	v.voice, err = dg.ChannelVoiceJoin(v.guildID, voiceChannelID, false, true)
	if err != nil {
		v.Stop()
		log.Println("Error when joining voice channel")
		return
	}

	v.voice.Speaking(false)
	log.Println("New voice instance created")
	s.ChannelMessageSend(m.ChannelID, "[Music] Voice channel joined!")
}


func leaveVoice(v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Println("INFO: ", m.Author.Username, "requested 'leave'")

	if v == nil {
		log.Println("INFO: The bot is not in a voice channel!")
		return
	}

	v.Stop()
	time.Sleep(200 * time.Millisecond)
	v.voice.Disconnect()
	log.Println("INFO: Voice channel left")
	mutex.Lock()
	delete(voiceInstances, v.guildID)
	mutex.Unlock()
	dg.ChannelMessageSend(m.ChannelID, "[Music] Voice channel left!")
}


func playMusic(n []string, v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) {

	// fmt.Println("n = ", n)

	if v == nil {
		log.Println("INFO: The bot is not in a voice channel")
		dg.ChannelMessageSend(m.ChannelID, "[Music] I need to join a voice channel first!")
		return
	}

	voiceChannelID := searchVoiceChannel(m.Author.ID)
	if v.voice.ChannelID != voiceChannelID {
		s.ChannelMessageSend(m.ChannelID, "[Music] <@" + m.Author.ID + "> you need to join my voice channel first!")
		return
	}

	name := replaceSpace(n)
	// url := youtubeEndpoint + conf.Ytkey + "&q=" + name

	song, err := ytSearch(name, v, s, m)
	if err != nil || song.data.ID == "" {
		log.Println("ERROR: Youtube search: ", err)
		s.ChannelMessageSend(m.ChannelID, "[Music] Unable to find song")
		return
	}

	s.ChannelMessageSend(m.ChannelID, "[Music] " + song.data.User + " has added " + song.data.Title + " to the queue")

	go func() {
		songSignal <- song
	}()


	// if result[0] == ytSearchFailed || result[0] == decodingFailed {
	// 	log.Println("ERROR: Something wrong happened when getting song")
	// 	s.ChannelMessageSend(m.ChannelID, "[Music] Something wrong happened when getting song")
	// 	return
	// }



	// vid, err := v2.GetVideoInfo(ytVideoUrl + result[0])
	// if err != nil {
	// 	return
	// }

	// client := youtube.Client{}
	// video, err := client.GetVideo(result[0])
	// if err != nil {
	// 	return
	// }

	// go func () {
	// 	songSignal <- video
	// }()


	// s.ChannelMessageSend(m.ChannelID, result[0] + "\t" + result[1])
}

func skipMusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Skip music!")
}
