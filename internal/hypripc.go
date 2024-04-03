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
	addr string
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

func (hypr HyprConn) SendNotification(nf *Notification) {

	icon := i32ToString(nf.icon.value)
	timeout := i32ToString(nf.time_ms)
	font_size := i32ToString(nf.font_size.value)
	msg := "fontsize:" + font_size + " " + nf.icon.padding + nf.message

	hypr.HyprCtl("notify", icon, timeout, nf.color.value, msg)
}

func (hypr HyprConn) DismissNotify(last int) {
	amount := strconv.Itoa(last)
	hypr.HyprCtl("dismissnotify", amount)
}

func GetHyprSocket(hypr *HyprConn) {
	hyprsock.addr = GetHyprSocketAddr()
}
