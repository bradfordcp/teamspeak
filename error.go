package teamspeak

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Error struct {
	Id  uint
	Msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Id, e.Msg)
}

// Parse a string and convert it to a Error
func NewError(errorStr string) (*Error, error) {
	ts3Err := &Error{}

	// Parse the error line
	tokens := strings.Split(string(errorStr), " ")
	if "error" == tokens[0] {
		for _, token := range tokens[1:] {
			attribute := strings.Split(token, "=")
			switch attribute[0] {
			case "id":
				id, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return ts3Err, err
				}
				ts3Err.Id = uint(id)
			case "msg":
				ts3Err.Msg = attribute[1]
			default:
				// We don't recognize the error, emit an error about the attribute
				return ts3Err, errors.New(fmt.Sprintf("Error could not parse param: %v", attribute[0]))
			}
		}
	} else {
		// First token was not error, emit an error about not parsing the error
		return ts3Err, errors.New(fmt.Sprintf("Error could not parse error from: %v", errorStr))
	}

	return ts3Err, nil
}
