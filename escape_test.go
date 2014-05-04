package teamspeak

import (
	"testing"
)

func escapeTestHelper(in, expected string, t *testing.T) {
	out := Escape(in)
	if out != expected {
		t.Errorf("Escape(%v) should have returned %v, not %v", in, expected, out)
	}
}

func unescapeTestHelper(in, expected string, t *testing.T) {
	out := Unescape(in)
	if out != expected {
		t.Errorf("Unescape(%v) should have returned %v, not %v", in, expected, out)
	}
}

func TestEscape(t *testing.T) {
	// Test a vanilla string with no escaped characters
	escapeTestHelper("foobarbaz", "foobarbaz", t)

	// Backslash!
	escapeTestHelper("\\", "\\\\", t)

	// Slash
	escapeTestHelper("/", "\\/", t)

	// Whitespace
	escapeTestHelper(" ", "\\s", t)

	// Pipe
	escapeTestHelper("|", "\\p", t)

	// Bell
	escapeTestHelper("\a", "\\a", t)

	// Backspace
	escapeTestHelper("\b", "\\b", t)

	// Formfeed
	escapeTestHelper("\f", "\\f", t)

	// Newline
	escapeTestHelper("\n", "\\n", t)

	// Carraige Return
	escapeTestHelper("\r", "\\r", t)

	// Horizontal Tab
	escapeTestHelper("\t", "\\t", t)

	// Vertical Tab
	escapeTestHelper("\v", "\\v", t)

	// All the things
	escapeTestHelper("foo\\/ |\a\b\f\n\r\t\v", "foo\\\\\\/\\s\\p\\a\\b\\f\\n\\r\\t\\v", t)
}

func TestUnescape(t *testing.T) {
	// Test a vanilla string with no escaped characters
	unescapeTestHelper("foobarbaz", "foobarbaz", t)

	// Backslash!
	unescapeTestHelper("\\\\", "\\", t)

	// Slash
	unescapeTestHelper("\\/", "/", t)

	// Whitespace
	unescapeTestHelper("\\s", " ", t)

	// Pipe
	unescapeTestHelper("\\p", "|", t)

	// Bell
	unescapeTestHelper("\\a", "\a", t)

	// Backspace
	unescapeTestHelper("\\b", "\b", t)

	// Formfeed
	unescapeTestHelper("\\f", "\f", t)

	// Newline
	unescapeTestHelper("\\n", "\n", t)

	// Carraige Return
	unescapeTestHelper("\\r", "\r", t)

	// Horizontal Tab
	unescapeTestHelper("\\t", "\t", t)

	// Vertical Tab
	unescapeTestHelper("\\v", "\v", t)

	// All the things
	unescapeTestHelper("foo\\\\\\/\\s\\p\\a\\b\\f\\n\\r\\t\\v", "foo\\/ |\a\b\f\n\r\t\v", t)
}
