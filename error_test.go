package teamspeak

import (
	"testing"
)

func TestNewError(t *testing.T) {
	const validErrorString = "error id=0 msg=ok"
	const invalidErrorString = "I'm not an error"
	const invalidErrorParamString = "error foo=bar"

	// Test to see if a valid error string is converted into a Channel struct
	validError, err := NewError(validErrorString)
	if err != nil {
		t.Errorf("NewError(\"%v\"): Errored out with %v", validErrorString, err)
	} else {
		if validError.Id != 0 || validError.Msg != "ok" {
			t.Errorf("NewError(\"%v\"): Parsed version %v does not match source input", validErrorString, validError)
		}
	}

	// Test to see if a invalid error string throws an error
	invalidError, err := NewError(invalidErrorString)
	if err == nil {
		t.Errorf("NewError(\"%v\"): Should have thrown an error. Instead received %v", invalidErrorString, invalidError)
	}

	// Test to see if a invalid error param string throws an error
	invalidParamError, err := NewError(invalidErrorParamString)
	if err == nil {
		t.Errorf("NewError(\"%v\"): Should have thrown an error. Instead received %v", invalidErrorParamString, invalidParamError)
	}
}
