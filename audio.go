package discordbot

import (
	"os/exec"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

type VoiceInstance struct {
	voice      *discordgo.VoiceConnection
	session    *discordgo.Session
	encoder    *dca.EncodeSession
	stream     *dca.StreamingSession
	run        *exec.Cmd
	queueMutex sync.Mutex
	audioMutex sync.Mutex
	nowPlaying Song
	queue      []Song
	recv       []int16
	guildID    string
	channelID  string
	speaking   bool
	pause      bool
	stop       bool
	skip       bool
	radioFlag  bool
}

type Song struct {
	ChannelID string
	User      string
	ID        string
	VidID     string
	Title     string
	Duration  string
	VideoURL  string
}

var (
	voiceInstances = map[string]*VoiceInstance{}
	mutex sync.Mutex
)

// Stop stops the audio
func (v *VoiceInstance) Stop() {
	v.stop = true
	if v.encoder != nil {
		v.encoder.Cleanup()
	}
}
