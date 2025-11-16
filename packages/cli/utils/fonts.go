package utils

import (
	"fmt"
	"strings"
)

// helper to convert int to string (alternative to strconv.Itoa for brevity)
func itoa(i int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(i)), ""), "[]")
}

// Color returns the input string with the given color code (30–37 for text).
func color(text string, colorCode int) string {
	return "\033[" + itoa(colorCode) + "m" + text + "\033[0m"
}

// Common shorthand for basic colors
func Green(text string) string  { return color(text, 32) }
func Yellow(text string) string { return color(text, 33) }
func Red(text string) string    { return color(text, 31) }
func Cyan(text string) string   { return color(text, 36) }
func White(text string) string  { return color(text, 37) }

// Weight

const ansiReset = "\033[0m"

func Bold(text string) string {
	return "\033[1m" + text + ansiReset
}

// Italic returns the input string in italic (may not work on all terminals).
func Italic(text string) string {
	return "\033[3m" + text + ansiReset
}

// Underline returns the input string underlined.
func Underline(text string) string {
	return "\033[4m" + text + ansiReset
}
