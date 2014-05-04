package teamspeak

import (
	"bytes"
)

func Escape(toEscape string) string {
	escapedBytes := []byte(toEscape)

	// Backslashes
	escapedBytes = bytes.Replace(escapedBytes, []byte{92}, []byte{92, 92}, -1)

	// Slash
	escapedBytes = bytes.Replace(escapedBytes, []byte{47}, []byte{92, 47}, -1)

	// Whitespace
	escapedBytes = bytes.Replace(escapedBytes, []byte{32}, []byte{92, 115}, -1)

	// Pipe
	escapedBytes = bytes.Replace(escapedBytes, []byte{124}, []byte{92, 112}, -1)

	// Bell
	escapedBytes = bytes.Replace(escapedBytes, []byte{7}, []byte{92, 97}, -1)

	// Backspace
	escapedBytes = bytes.Replace(escapedBytes, []byte{8}, []byte{92, 98}, -1)

	// Formfeed
	escapedBytes = bytes.Replace(escapedBytes, []byte{12}, []byte{92, 102}, -1)

	// Newline
	escapedBytes = bytes.Replace(escapedBytes, []byte{10}, []byte{92, 110}, -1)

	// Carraige Return
	escapedBytes = bytes.Replace(escapedBytes, []byte{13}, []byte{92, 114}, -1)

	// Horizontal Tab
	escapedBytes = bytes.Replace(escapedBytes, []byte{9}, []byte{92, 116}, -1)

	// Vertical Tab
	escapedBytes = bytes.Replace(escapedBytes, []byte{11}, []byte{92, 118}, -1)

	return string(escapedBytes)
}

func Unescape(toUnescape string) string {
	unescapedBytes := []byte(toUnescape)

	// Backslashes
	unescapedBytes = bytes.Replace(unescapedBytes, []byte{92, 92}, []byte{92}, -1)

	// Slash
	unescapedBytes = bytes.Replace(unescapedBytes, []byte{92, 47}, []byte{47}, -1)

	// Whitespace
	unescapedBytes = bytes.Replace(unescapedBytes, []byte{92, 115}, []byte{32}, -1)

	// Pipe
	unescapedBytes = bytes.Replace(unescapedBytes, []byte{92, 112}, []byte{124}, -1)

	// Bell
	unescapedBytes = bytes.Replace(unescapedBytes, []byte{92, 97}, []byte{7}, -1)

	// Backspace
	unescapedBytes = bytes.Replace(unescapedBytes, []byte{92, 98}, []byte{8}, -1)

	// Formfeed
	unescapedBytes = bytes.Replace(unescapedBytes, []byte{92, 102}, []byte{12}, -1)

	// Newline
	unescapedBytes = bytes.Replace(unescapedBytes, []byte{92, 110}, []byte{10}, -1)

	// Carraige Return
	unescapedBytes = bytes.Replace(unescapedBytes, []byte{92, 114}, []byte{13}, -1)

	// Horizontal Tab
	unescapedBytes = bytes.Replace(unescapedBytes, []byte{92, 116}, []byte{9}, -1)

	// Vertical Tab
	unescapedBytes = bytes.Replace(unescapedBytes, []byte{92, 118}, []byte{11}, -1)

	return string(unescapedBytes)
}
