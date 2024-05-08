package internal

import (
	"fmt"
	"net"
	"os"
	"path"
	"strconv"
	"strings"
)

func GetHyprSocketAddr() string {
	instance_signature := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	if instance_signature == "" {
		fmt.Println("Hyprland is not running!")
		os.Exit(1)
	}

	runtime_dir := os.Getenv("XDG_RUNTIME_DIR")
	if runtime_dir == "" {
		runtime_dir = "/run/user/1000" // try for first user
	}

	socket_addr := path.Join(runtime_dir, "/hypr", instance_signature, "/.socket.sock")

	if _, err := os.Stat(socket_addr); err == nil {
		return socket_addr
	}

	// try pre v40 socket path
	socket_addr = path.Join("/tmp/hypr/", instance_signature, "/.socket.sock")

	if _, err := os.Stat(socket_addr); err == nil {
		return socket_addr
	}

	fmt.Println("Hyprland IPC path is not available!")
	os.Exit(1)
	return ""
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
