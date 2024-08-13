package color

import "regexp"

// GetRawString 返回没有颜色转译字符的原始字符串
func GetRawString(colorString string) string {
	ansiRegex := regexp.MustCompile("\033\\[[0-9;]+m")
	rawString := ansiRegex.ReplaceAllString(colorString, "")
	return rawString
}
