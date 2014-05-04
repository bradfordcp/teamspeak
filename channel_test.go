package teamspeak

import (
	"testing"
)

const validChannelPropertyString = "cid=1 pid=2 channel_order=3 channel_name=Sample\\sChannel\\sName total_clients=4 channel_needed_subscribe_power=5"
const validChannelPropertyStringWithNull = "cid=1 pid channel_order=3 channel_name=Sample\\sChannel\\sName total_clients=4 channel_needed_subscribe_power=5"
const invalidChannelPropertyString = "I'm not a channel, nice try!"

func TestNewChannel(t *testing.T) {
	// Test to see if a valid channel string is converted into a Channel struct
	validChannel, err := NewChannel(validChannelPropertyString)
	if err != nil {
		t.Errorf("NewChannel(\"%v\"): Errored out with %v", validChannelPropertyString, err)
	} else {
		if validChannel.Cid != 1 || validChannel.Pid != 2 || validChannel.ChannelName != "Sample\\sChannel\\sName" || validChannel.TotalClients != 4 || validChannel.ChannelNeededSubscribePower != 5 {
			t.Errorf("NewChannel(\"%v\"): Parsed version %v does not match source input", validChannelPropertyString, validChannel)
		}
	}

	// Test to see if a valid channel string is converted into a Channel struct, (with null values)
	validNullParamChannel, err := NewChannel(validChannelPropertyStringWithNull)
	if err != nil {
		t.Errorf("NewChannel(\"%v\"): Errored out with %v", validChannelPropertyStringWithNull, err)
	} else {
		if validNullParamChannel.Cid != 1 || validNullParamChannel.Pid != 0 || validNullParamChannel.ChannelName != "Sample\\sChannel\\sName" || validNullParamChannel.TotalClients != 4 || validNullParamChannel.ChannelNeededSubscribePower != 5 {
			t.Errorf("NewChannel(\"%v\"): Parsed version %v does not match source input", validChannelPropertyStringWithNull, validNullParamChannel)
		}
	}

	// Test to see if a invalid channel string throws an error
	invalidChannel, err := NewChannel(invalidChannelPropertyString)
	if err == nil {
		t.Errorf("NewChannel(\"%v\"): Should have thrown an error. Instead received %v", invalidChannelPropertyString, invalidChannel)
	}
}

func TestDeserialize(t *testing.T) {
	const validChannelUpdateString = "cid=2 pid=3"
	const invalidChannelUpdateString = "cid=4 invalid=true"

	// Test to see if a valid channel string is converted into a Channel struct
	validChannel, err := NewChannel(validChannelPropertyString)
	if err != nil {
		t.Errorf("NewChannel(\"%v\"): Errored out with %v", validChannelPropertyString, err)
	}

	// Test a valid update call
	_, err = validChannel.Deserialize(validChannelUpdateString)

	if err != nil {
		t.Errorf("channel.Deserialize(\"%v\"): Errored out with %v", validChannelUpdateString, err)
	}

	// Test an invalid update call
	_, err = validChannel.Deserialize(invalidChannelUpdateString)

	if err == nil {
		t.Errorf("channel.Deserialize(\"%v\"): should have thrown an error", invalidChannelUpdateString, err)
	}
}
