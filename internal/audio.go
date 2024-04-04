package internal

import (
	"time"
  "embed"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

//go:embed audio
var fs embed.FS

func PlayAudio() {
	f, err := fs.Open("audio/notify_beep.wav")
	if err != nil {
		panic(err)
	}
	streamer, _, err := wav.Decode(f)
	if err != nil {
		panic(err)
	}
	defer streamer.Close()
	speaker.Play(streamer)
	for streamer.Len() != streamer.Position() {
		time.Sleep(time.Second)
	}
}

func InitSpeaker() {
	var sr beep.SampleRate = 44100
	speaker.Init(sr, sr.N(time.Second/10))

}
