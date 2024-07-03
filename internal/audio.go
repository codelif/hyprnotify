package internal

import (
	"embed"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

//go:embed audio
var fs embed.FS

func PlayAudio() {
	f, err := fs.Open("audio/notify_beep.wav")
	if err != nil {
		panic(err)
	}

	ctx := audio.CurrentContext()

	player, err := ctx.NewPlayer(f)
	if err != nil {
		panic(err)
	}
	defer player.Close()

	player.Play()
	for player.IsPlaying() {
		time.Sleep(time.Second)
	}
}

func InitSpeaker() {
	// var sr beep.SampleRate = 44100
	// speaker.Init(sr, sr.N(time.Second/10))
	audio.NewContext(44100)
}
