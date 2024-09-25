package main

import (
	"github.com/codelif/hyprnotify/internal"
	"github.com/spf13/cobra"
)

func main() {
	var enableSound bool

	Cmd := &cobra.Command{
		Use:  "hyprnotify",
		Long: `DBus Implementation of Freedesktop Notification spec for 'hyprctl notify'`,
		Run: func(cmd *cobra.Command, args []string) {
			internal.InitDBus(enableSound)
		},
	}

	CmdFlags := Cmd.Flags()

	CmdFlags.BoolVarP(&enableSound, "no-sound", "s", false, "disable sound, silent mode")

	Cmd.Execute()
}
