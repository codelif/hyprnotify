package main

import (
	"flag"
	"github.com/codelif/hyprnotify/internal"
)

func main() {
	var enableSound bool
	const message = "Disable sound"

	flag.BoolVar(&enableSound, "no-sound", false, message)
	flag.BoolVar(&enableSound, "silent", false, message)
	flag.BoolVar(&enableSound, "s", false, message)

	flag.Parse()

	internal.InitDBus(enableSound)
}
