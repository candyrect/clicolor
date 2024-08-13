package color

import "testing"

func TestColor(t *testing.T) {
	t.Log("Testing color package")
	LightGreen.Underline().Blink().Bold().Println("Hello, World!")
}
