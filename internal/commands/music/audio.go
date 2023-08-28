package music

import (
	"github.com/Dahl99/discord-bot/internal/discord"
	"io"
	"log"
	"log/slog"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/jung-m/dca"
)

type VoiceInstance struct {
	voice      *discordgo.VoiceConnection
	session    *discordgo.Session
	encoder    *dca.EncodeSession
	stream     *dca.StreamingSession
	queueMutex sync.Mutex
	nowPlaying Song
	queue      []Song
	GuildID    string
	speaking   bool
	pause      bool
	stop       bool
	skip       bool
}

type Song struct {
	ChannelID string
	User      string
	ID        string
	VidID     string
	Title     string
	VideoURL  string
}

type PkgSong struct {
	data Song
	v    *VoiceInstance
}

var (
	VoiceInstances = map[string]*VoiceInstance{}
	mutex          sync.Mutex
	SongSignal     chan PkgSong
)

func globalPlay(songSig chan PkgSong) {
	for {
		select {
		case song := <-songSig:
			go song.v.PlayQueue(song.data)
		}
	}
}

func (v *VoiceInstance) Skip() bool {
	if v.speaking {
		if v.pause {
			return true
		} else {
			if v.encoder != nil {
				v.encoder.Cleanup()
			}
		}
	}
	return false
}

// Stop stops the audio
func (v *VoiceInstance) Stop() {
	v.stop = true
	if v.encoder != nil {
		v.encoder.Cleanup()
	}
}

// QueueAdd
func (v *VoiceInstance) QueueAdd(song Song) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	v.queue = append(v.queue, song)
}

// QueueGetSong
func (v *VoiceInstance) QueueGetSong() (song Song) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 {
		return v.queue[0]
	}
	return
}

// QueueRemoveFirst
func (v *VoiceInstance) QueueRemoveFirst() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 {
		v.queue = v.queue[1:]
	}
}

// QueueRemove
func (v *VoiceInstance) QueueRemove() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	v.queue = []Song{}
}

// DCA
func (v *VoiceInstance) DCA(url string) {
	opts := dca.StdEncodeOptions
	opts.RawOutput = true
	opts.Bitrate = 64
	opts.Application = "lowdelay"

	encodeSession, err := dca.EncodeFile(url, opts)
	if err != nil {
		slog.Error("Failed to create an encoding session", "error", err)
	}

	v.encoder = encodeSession
	done := make(chan error)
	v.stream = dca.NewStream(encodeSession, v.voice, done)
	for {
		select {
		case err := <-done:
			if err != nil && err != io.EOF {
				slog.Error("An error occured", "error", err)
			}

			// Clean up in case something happened and ffmpeg is still running
			encodeSession.Cleanup()
			return
		}
	}
}

func (v *VoiceInstance) PlayQueue(song Song) {
	// add song to queue
	v.QueueAdd(song)

	if v.speaking {
		// the bot is playing
		return
	}

	go func() {
		for {
			if len(v.queue) == 0 {
				log.Println("INFO: End of queue")
				discord.SendChannelMessage(v.nowPlaying.ChannelID, "**[Music]** End of queue")
				return
			}

			v.nowPlaying = v.QueueGetSong()
			go discord.SendChannelMessage(v.nowPlaying.ChannelID, "**[Music]** Now playing: **"+v.nowPlaying.Title+"**")

			v.stop = false
			v.skip = false
			v.speaking = true
			v.pause = false
			err := v.voice.Speaking(true)
			if err != nil {
				slog.Error("Failed to send speaking notification", "error", err)
				return
			}

			v.DCA(v.nowPlaying.VideoURL)

			v.QueueRemoveFirst()
			if v.stop {
				v.QueueRemove()
			}

			v.stop = false
			v.skip = false
			v.speaking = false

			err = v.voice.Speaking(false)
			if err != nil {
				slog.Error("Failed to stop sending speaking notification", "error", err)
			}
		}
	}()
}
