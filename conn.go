package teamspeak

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
)

type Conn struct {
	conn  *net.TCPConn
	rw    *bufio.ReadWriter
	Debug bool
}

// Generates a new connection, dials out, and verifies connectivity
func NewConn(connectionString string) (*Conn, error) {
	// Resolve the address we are connecting to
	tsAddr, err := net.ResolveTCPAddr("tcp", connectionString)
	if err != nil {
		return nil, err
	}

	// Set up the object to return
	ts3 := &Conn{}

	// Dial the remote address
	ts3.conn, err = net.DialTCP("tcp", nil, tsAddr)
	if err != nil {
		return nil, err
	}

	// Setup the reader and writer
	reader := bufio.NewReader(ts3.conn)
	writer := bufio.NewWriter(ts3.conn)
	ts3.rw = bufio.NewReadWriter(reader, writer)

	// Read the first line and verify we are indeed connected to a TS server
	line, prefix, err := ts3.rw.ReadLine()
	if err != nil {
		return nil, err
	}
	if false == prefix && "TS3" != string(line) {
		ts3.conn.Close()
		return nil, errors.New("Not connected to a TS3 server")
	}

	// Read the next line, it is just help info
	ts3.rw.ReadLine()

	// Return the connection
	return ts3, nil
}

// Sends the command, which must already be encoded
func (ts3 *Conn) SendCommand(command string) (string, error) {
	if ts3.Debug {
		fmt.Println("SEND:", command)
	}

	// Send the command up with a newline added
	_, err := ts3.rw.WriteString(command + "\n")
	if err != nil {
		return "", err
	}

	// Flush the writer
	ts3.rw.Flush()

	// Return the response
	return ts3.ReadResponse()
}

func (ts3 *Conn) ReadResponse() (string, error) {
	// Generate the response data structure
	responseBuffer := make([]byte, 0)
	var ts3Err *Error

	continueReadingResponse := true
	for continueReadingResponse {
		lineBuffer := make([]byte, 0)

		// Read the response
		for continueReadingLine := true; continueReadingLine; {
			rawResponse, isPrefix, err := ts3.rw.ReadLine()
			if err != nil {
				return "", err
			}
			lineBuffer = append(lineBuffer, rawResponse...)

			continueReadingLine = isPrefix
		}

		if "error" == string(lineBuffer[1:6]) {
			// Last line of response has been detected
			continueReadingResponse = false
			var err error

			ts3Err, err = NewError(string(lineBuffer[1:]))
			if err != nil {
				return "", err
			}
		} else {
			// Store the text of the response and continue reading (next line will be error related)
			responseBuffer = append(responseBuffer, lineBuffer...)
		}
	}

	if ts3.Debug {
		fmt.Println("RECV:", string(responseBuffer), ts3Err)
	}

	return string(responseBuffer), ts3Err
}

// Authenticates with the username and password provided
func (ts3 *Conn) Login(username, password string) error {
	_, err := ts3.SendCommand("login " + username + " " + password)
	if ts3Err, ok := err.(*Error); ok && ts3Err.Id == 0 {
		return nil
	}

	return err
}

// Selects the virtual server to act on
func (ts3 *Conn) Use(serverId int) error {
	_, err := ts3.SendCommand("use sid=" + strconv.Itoa(serverId))
	if ts3Err, ok := err.(*Error); ok && ts3Err.Id == 0 {
		return nil
	}

	return err
}

// Closes the TCP Connection
func (ts3 *Conn) Close() {
	ts3.conn.Close()
}
