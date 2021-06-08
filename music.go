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
		leaveVoice(v, m)
	case "play":
		playMusic(cmd[1:], v, m)
	case "skip":
		skipMusic(v, m)
	default:
		return
	}
}

func joinVoice(v *VoiceInstance, s *discordgo.Session, m *discordgo.MessageCreate) {
	voiceChannelID := searchVoiceChannel(m.Author.ID)
	if voiceChannelID == "" {
		log.Println("Voice channel id not found")
		dg.ChannelMessageSend(m.ChannelID, "[Music] You need to join a voice channel first!")
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
	dg.ChannelMessageSend(m.ChannelID, "[Music] Voice channel joined!")
}


func leaveVoice(v *VoiceInstance, m *discordgo.MessageCreate) {
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


func playMusic(n []string, v *VoiceInstance, m *discordgo.MessageCreate) {
	if v == nil {
		log.Println("INFO: The bot is not in a voice channel")
		dg.ChannelMessageSend(m.ChannelID, "[Music] I need to join a voice channel first!")
		return
	}

	voiceChannelID := searchVoiceChannel(m.Author.ID)
	if v.voice.ChannelID != voiceChannelID {
		dg.ChannelMessageSend(m.ChannelID, "[Music] <@" + m.Author.ID + "> you need to join my voice channel first!")
		return
	}

	name := replaceSpace(n)
	videoId, videoTitle, err := ytSearch(name)
	if err != nil {
		log.Println("INFO: Unable to find song by searching youtube")
		dg.ChannelMessageSend(m.ChannelID, "[Music] Oops, something wrong happened when searching YouTube")
		return
	}

	song, err := execYtdl(videoId, videoTitle, v, m)

	if err != nil || song.data.ID == "" {
		log.Println("ERROR: Youtube search: ", err)
		dg.ChannelMessageSend(m.ChannelID, "[Music] Unable to find song")
		return
	}

	dg.ChannelMessageSend(m.ChannelID, "[Music] " + song.data.User + " has added " + song.data.Title + " to the queue")

	go func() {
		songSignal <- song
	}()
}

func skipMusic(v *VoiceInstance, m *discordgo.MessageCreate) {
	log.Println("INFO:", m.Author.Username, "send 'skip'")
	if v == nil {
		log.Println("INFO: The bot is not in a voice channel")
		dg.ChannelMessageSend(m.ChannelID, "[Music] I need to join a voice channel first!")
		return
	}
	if len(v.queue) == 0 {
		log.Println("INFO: The queue is empty.")
		dg.ChannelMessageSend(m.ChannelID, "[Music] There is no song playing")
		return
	}
	if v.Skip() {
		dg.ChannelMessageSend(m.ChannelID, "[Music] I'm paused, resume first")
	}
}
