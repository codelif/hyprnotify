package internal

type Notification struct {
	message string

	time_ms   int32
	icon      icon
	color     color
	font_size fontSize
}

type color struct {
	value string

	DEFAULT   string
	VIOLET    string
	INDIGO    string
	BLUE      string
	LIGHTBLUE string
	GREEN     string
	YELLOW    string
	ORANGE    string
	RED       string
}

func (c *color) HEX(hexcode string) string {
	if string(hexcode[0]) == "#" {
		hexcode = hexcode[1:]
	}
	return "rgb(" + hexcode + ")"
}

type icon struct {
	value   int32
	padding string

	NOICON   int32
	WARNING  int32
	INFO     int32
	HINT     int32
	ERROR    int32
	CONFUSED int32
	OK       int32
}

type fontSize struct {
	value   int32
	DEFAULT int32
}

func newColorStruct() color {
	color := color{}

	color.DEFAULT = "0"
	color.VIOLET = color.HEX("9400D3")
	color.INDIGO = color.HEX("4B0082")
	color.BLUE = color.HEX("0000FF")
	color.LIGHTBLUE = color.HEX("00d2ff")
	color.GREEN = color.HEX("00FF00")
	color.YELLOW = color.HEX("FFFF00")
	color.ORANGE = color.HEX("FF7F00")
	color.RED = color.HEX("FF0000")

	return color
}

func newIconStruct() icon {
	icon := icon{}

	icon.NOICON = -1
	icon.WARNING = 0
	icon.INFO = 1
	icon.HINT = 2
	icon.ERROR = 3
	icon.CONFUSED = 4
	icon.OK = 5

	return icon
}

func (nf *Notification) set_urgency(urgency uint8) {
	icon := nf.icon.NOICON
	padding := ""
	color := nf.color.DEFAULT
	var time_ms int32 = 5 * 1000

	if urgency == 0 {
		icon = nf.icon.OK
		padding = " "
		color = nf.color.GREEN
	} else if urgency == 1 {
		icon = nf.icon.NOICON
		padding = " "
		color = nf.color.LIGHTBLUE
	} else if urgency == 2 {
		icon = nf.icon.WARNING
		padding = "  "
		color = nf.color.RED
		time_ms = 60 * 1000
	}

	nf.icon.value = icon
	nf.icon.padding = padding
	nf.color.value = color
	nf.time_ms = time_ms
}

func NewNotification() Notification {
	n := Notification{}

	n.icon = newIconStruct()
	n.color = newColorStruct()
	n.font_size = fontSize{value: 13, DEFAULT: 13}

	n.set_urgency(1) // default
	return n
}
