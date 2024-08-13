package color

// Bold 加粗
func (c Color) Bold() Color {
	return Color("\033[1m" + string(c))
}

// Faint 弱化颜色
func (c Color) Faint() Color {
	return Color("\033[2m" + string(c))
}

// Italic 斜体
func (c Color) Italic() Color {
	return Color("\033[3m" + string(c))
}

// Underline 下划线
func (c Color) Underline() Color {
	return Color("\033[4m" + string(c))
}

// Blink 闪烁
func (c Color) Blink() Color {
	return Color("\033[5m" + string(c))
}

// Invert 反转颜色
func (c Color) Invert() Color {
	return Color("\033[7m" + string(c))
}

// Strikethrough 删除线
func (c Color) Strikethrough() Color {
	return Color("\033[9m" + string(c))
}
