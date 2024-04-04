package main

import (
	"github.com/codelif/hyprnotify/internal"
	"os"
)

func main() {
	var enable_sound bool = true

	for _, arg := range os.Args[1:] {
		if arg == "--no-sound" || arg == "--silent" || arg == "-s" {
			enable_sound = false
			break
		}
	}

	internal.InitDBus(enable_sound)
}
