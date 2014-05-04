package teamspeak

import (
	"testing"
)

func TestNewChannel(t *testing.T) {
	const validChannelString = "cid=1 pid=2 channel_order=3 channel_name=Sample\\sChannel\\sName total_clients=4 channel_needed_subscribe_power=5"
	const invalidChannelString = "I'm not a channel, nice try!"

	// Test to see if a valid channel string is converted into a Channel struct
	validChannel, err := NewChannel(validChannelString)
	if err != nil {
		t.Errorf("NewChannel(\"%v\"): Errored out with %v", validChannelString, err)
	} else {
		if validChannel.Cid != 1 || validChannel.Pid != 2 || validChannel.ChannelName != "Sample\\sChannel\\sName" || validChannel.TotalClients != 4 || validChannel.ChannelNeededSubscribePower != 5 {
			t.Errorf("NewChannel(\"%v\"): Parsed version %v does not match source input", validChannelString, validChannel)
		}
	}

	// Test to see if a invalid channel string throws an error
	invalidChannel, err := NewChannel(invalidChannelString)
	if err == nil {
		t.Errorf("NewChannel(\"%v\"): Should have thrown an error. Instead received %v", invalidChannelString, invalidChannel)
	}
}
