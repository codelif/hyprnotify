package internal

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func GetHyprSocketAddr() string {
	env := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	if env == "" {
		fmt.Println("Hyprland is not running!")
		os.Exit(1)
	}

	addr := "/tmp/hypr/" + env + "/.socket.sock"

	return addr
}

type HyprConn struct {
	addr              string
	icon              Icons
	color             Colors
	DEFAULT_FONT_SIZE string
}

type Colors struct {
	DEFAULT string
	VIOLET  string
	INDIGO  string
	BLUE    string
	GREEN   string
	YELLOW  string
	ORANGE  string
	RED     string
}

type Icons struct {
	NOICON   string
	WARNING  string
	INFO     string
	HINT     string
	ERROR    string
	CONFUSED string
	OK       string
}

func (c Colors) HEX(hexcode string) string {
	if string(hexcode[0]) == "#" {
		hexcode = hexcode[1:]
	}
	return "rgb(" + hexcode + ")"
}

func (hypr HyprConn) GetConn() {
}

func (hypr HyprConn) HyprCtl(args ...string) {
	conn, err := net.Dial("unix", hyprsock.addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	msg := "/" + strings.Join(args, " ")

	_, err = conn.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
}

func (hypr HyprConn) Notify(
	icon string,
	time_ms int32,
	color string,
	msg string,
	font_size string,
) {
	fmt_msg := "fontsize:" + font_size + " " + msg
	fmt_time := strconv.FormatInt(int64(time_ms), 10)
	hypr.HyprCtl("notify", icon, fmt_time, color, fmt_msg)
}

func (hypr HyprConn) DismissNotify(last int) {
	amount := strconv.Itoa(last)
	hypr.HyprCtl("dismissnotify", amount)
}

func attachIconsStruct(hypr *HyprConn) {
	var icons Icons
	icons.NOICON = "-1"
	icons.WARNING = "0"
	icons.INFO = "1"
	icons.HINT = "2"
	icons.ERROR = "3"
	icons.CONFUSED = "4"
	icons.OK = "5"

	hypr.icon = icons
}

func attachColorsStruct(hypr *HyprConn) {
	var color Colors
	color.DEFAULT = "0"
	color.VIOLET = color.HEX("9400D3")
	color.INDIGO = color.HEX("4B0082")
	color.BLUE = color.HEX("0000FF")
	color.GREEN = color.HEX("00FF00")
	color.YELLOW = color.HEX("FFFF00")
	color.ORANGE = color.HEX("FF7F00")
	color.RED = color.HEX("FF0000")

	hypr.color = color
}

func GetHyprSocket(hypr *HyprConn) {
	hyprsock.addr = GetHyprSocketAddr()
	attachIconsStruct(hypr)
	attachColorsStruct(hypr)
	hypr.DEFAULT_FONT_SIZE = "13"

	hyprsock.Notify(
		hyprsock.icon.INFO,
		10000,
		hyprsock.color.GREEN,
		"Test",
		hyprsock.DEFAULT_FONT_SIZE,
	)
}
