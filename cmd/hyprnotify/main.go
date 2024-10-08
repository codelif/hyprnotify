package main

import (
	"github.com/codelif/hyprnotify/internal"
	"github.com/spf13/cobra"
)

func main() {
	var disableSound bool

	Cmd := &cobra.Command{
		Use:  "hyprnotify",
		Long: `DBus Implementation of Freedesktop Notification spec for 'hyprctl notify'`,
		Run: func(cmd *cobra.Command, args []string) {
			internal.InitDBus(!disableSound)
		},
	}

	CmdFlags := Cmd.Flags()

	CmdFlags.BoolVarP(&disableSound, "no-sound", "s", false, "disable sound, silent mode")
	CmdFlags.Uint8VarP(&internal.DefaultFontSize, "font-size", "f", 13, "set default font size (range 1-255)")
	CmdFlags.BoolVar(&internal.FixedFontSize, "fixed-font-size", false, "makes font size fixed, ignoring new sizes")

	Cmd.Execute()
}
