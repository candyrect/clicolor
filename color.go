package color

import (
	"fmt"
	"strings"
)

type Color string

const (
	Reset Color = "\033[0m"

	Green   Color = "\033[32m"
	Red     Color = "\033[31m"
	Blue    Color = "\033[34m"
	Cyan    Color = "\033[36m"
	Magenta Color = "\033[35m"
	Yellow  Color = "\033[33m"
	Grey    Color = "\033[37m"

	LightBlack   Color = "\033[90m"
	LightRed     Color = "\033[91m"
	LightGreen   Color = "\033[92m"
	LightYellow  Color = "\033[93m"
	LightBlue    Color = "\033[94m"
	LightMagenta Color = "\033[95m"
	LightCyan    Color = "\033[96m"
	LightWhite   Color = "\033[97m"
	LightGrey    Color = "\033[107m"

	BgBlack   Color = "\033[40m"
	BgRed     Color = "\033[41m"
	BgGreen   Color = "\033[42m"
	BgYellow  Color = "\033[43m"
	BgBlue    Color = "\033[44m"
	BgMagenta Color = "\033[45m"
	BgCyan    Color = "\033[46m"
	BgWhite   Color = "\033[47m"
	BgGrey    Color = "\033[100m"

	BgLightBlack   Color = "\033[100m"
	BgLightRed     Color = "\033[101m"
	BgLightGreen   Color = "\033[102m"
	BgLightYellow  Color = "\033[103m"
	BgLightBlue    Color = "\033[104m"
	BgLightMagenta Color = "\033[105m"
	BgLightCyan    Color = "\033[106m"
	BgLightWhite   Color = "\033[107m"
	BgLightGrey    Color = "\033[108m"
)

func (c Color) Sprint(a ...interface{}) string {
	sb := strings.Builder{}
	sb.WriteString(string(c))
	sb.WriteString(fmt.Sprint(a...))
	sb.WriteString(string(Reset))
	return sb.String()
}

func (c Color) Sprintf(format string, a ...interface{}) string {
	sb := strings.Builder{}
	sb.WriteString(string(c))
	sb.WriteString(fmt.Sprintf(format, a...))
	sb.WriteString(string(Reset))
	return sb.String()
}

func (c Color) Println(a ...interface{}) {
	fmt.Println(c.Sprint(a...))
}

func (c Color) Print(a ...interface{}) {
	fmt.Print(c.Sprint(a...))
}

func (c Color) GetCode() string {
	return string(c)
}
