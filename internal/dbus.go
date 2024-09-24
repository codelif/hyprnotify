package internal

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

const (
	FDN_PATH            = "/org/freedesktop/Notifications"
	FDN_IFAC            = "org.freedesktop.Notifications"
	FDN_NAME            = "org.freedesktop.Notifications"
	INTROSPECTABLE_IFAC = "org.freedesktop.DBus.Introspectable"

	FDN_SPEC_VERSION = "1.2"
	MAX_UINT32       = ^uint32(0)
)

const DBUS_XML = `<node name="` + FDN_PATH + `">
  <interface name="` + FDN_IFAC + `">

      <method name="GetCapabilities">
          <arg direction="out" name="capabilities"    type="as" />
      </method>

      <method name="Notify">
          <arg direction="in"  name="app_name"        type="s"/>
          <arg direction="in"  name="replaces_id"     type="u"/>
          <arg direction="in"  name="app_icon"        type="s"/>
          <arg direction="in"  name="summary"         type="s"/>
          <arg direction="in"  name="body"            type="s"/>
          <arg direction="in"  name="actions"         type="as"/>
          <arg direction="in"  name="hints"           type="a{sv}"/>
          <arg direction="in"  name="expire_timeout"  type="i"/>
          <arg direction="out" name="id"              type="u"/>
      </method> 

      <method name="GetServerInformation">
          <arg direction="out" name="name"            type="s"/>
          <arg direction="out" name="vendor"          type="s"/>
          <arg direction="out" name="version"         type="s"/>
          <arg direction="out" name="spec_version"    type="s"/>
      </method>

      <method name="CloseNotification">
          <arg direction="in"  name="id"              type="u"/>
      </method>

     <signal name="NotificationClosed">
          <arg name="id"         type="u"/>
          <arg name="reason"     type="u"/>
      </signal>

      <signal name="ActionInvoked">
          <arg name="id"         type="u"/>
          <arg name="action_key" type="s"/>
      </signal>
  </interface>
` + introspect.IntrospectDataString + `
</node>`

var (
	conn                  *dbus.Conn
	hyprsock              HyprConn
	ongoing_notifications map[uint32]chan uint32 = make(map[uint32]chan uint32)
	current_id            uint32                 = 0
	sound                 bool
)

type DBusNotify string

func (n DBusNotify) GetCapabilities() ([]string, *dbus.Error) {
	var cap []string
	return cap, nil
}

func (n DBusNotify) Notify(
	app_name string,
	replaces_id uint32,
	app_icon string,
	summary string,
	body string,
	actions []string,
	hints map[string]dbus.Variant,
	expire_timeout int32,
) (uint32, *dbus.Error) {

	if replaces_id > 0 {
		n.CloseNotification(replaces_id)
	}
	if current_id == MAX_UINT32 {
		current_id++
	}
	current_id++

	// Send Notification
	nf := NewNotification()
	if body != "" {
		nf.message = fmt.Sprintf("%s\n%s", summary, body)
	} else {
		nf.message = summary
	}

	// Using RegExp to add padding for all lines
	nf.message = regexp.
		MustCompile("^\\s*|(\n)\\s*(.)").
		ReplaceAllString(
			strings.TrimLeft(nf.message, "\n"),
			"$1\u205F\u205F$2",
		)

	parse_hints(&nf, hints)

	if expire_timeout != -1 {
		nf.time_ms = expire_timeout
	}
	hyprsock.SendNotification(&nf)

	if sound {
		go PlayAudio()
	}
	// ClosedNotification Signal Stuff
	flag := make(chan uint32, 1)
	ongoing_notifications[current_id] = flag
	go SendCloseSignal(nf.time_ms, current_id, 1, flag)
	return current_id, nil
}

func (n DBusNotify) CloseNotification(id uint32) *dbus.Error {
	count := 0
	for i := current_id; i >= id; i-- {
		flag, ok := ongoing_notifications[i]
		if ok {
			flag <- 3
		}
		count++
	}

	hyprsock.DismissNotify(count)

	return nil
}

func (n DBusNotify) GetServerInformation() (string, string, string, string, *dbus.Error) {
	return PACKAGE, VENDOR, VERSION, FDN_SPEC_VERSION, nil
}

func SendCloseSignal(timeout int32, id uint32, reason uint32, flag chan uint32) {
	d := time.Duration(int64(timeout)) * time.Millisecond

	tick := time.NewTicker(d)
	defer tick.Stop()

	select {
	case <-tick.C:
	case reason = <-flag:
	}
	conn.Emit(
		FDN_PATH,
		"org.freedesktop.Notifications.NotificationClosed",
		id,
		reason,
	)

	delete(ongoing_notifications, id)
}

func parse_hints(nf *Notification, hints map[string]dbus.Variant) {

	urgency, ok := hints["urgency"].Value().(uint8)
	if ok {
		nf.set_urgency(urgency)
	}

	font_size, ok := hints["x-hyprnotify-font-size"].Value().(int32)
	if ok {
		nf.font_size.value = font_size
	}
  
	hint_icon, ok := hints["x-hyprnotify-icon"].Value().(int32)
	if ok {
		nf.icon.value = hint_icon
		nf.icon.padding = ""
    nf.color.value = nf.color.DEFAULT
	}

	hint_color, ok := hints["x-hyprnotify-color"].Value().(string)
	if ok {
		if string(hint_color[0]) == "#" {
			hint_color = hint_color[1:]
		}
		if len(hint_color) == 6 && is_valid_hex_string(hint_color) {
			nf.color.value = nf.color.HEX(hint_color)
		}
	}


}

func InitDBus(enable_sound bool) {
	sound = enable_sound

	var err error
	conn, err = dbus.ConnectSessionBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	GetHyprSocket(&hyprsock)
	if sound {
		InitSpeaker()
	}

	n := DBusNotify(PACKAGE)
	conn.Export(n, FDN_PATH, FDN_IFAC)
	conn.Export(introspect.Introspectable(DBUS_XML), FDN_PATH, INTROSPECTABLE_IFAC)

	reply, err := conn.RequestName(FDN_NAME, dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}

	if reply != dbus.RequestNameReplyPrimaryOwner {
		fmt.Fprintln(os.Stderr, "name already taken")
		os.Exit(1)
	}
	select {}
}
