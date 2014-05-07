package teamspeak

import (
	"fmt"
	"strings"
	"testing"
)

const validChannelPropertyString = "cid=1 pid=2 channel_order=3 channel_name=Sample\\sChannel\\sName total_clients=4 channel_needed_subscribe_power=5"
const validChannelPropertyStringWithNull = "cid=1 pid channel_order=3 channel_name=Sample\\sChannel\\sName total_clients=4 channel_needed_subscribe_power=5"
const invalidChannelPropertyString = "invalid_property=foo"
const newChannelPropertyString = "pid=2 channel_order=5 channel_name=Woot"

func TestNewChannel(t *testing.T) {
	// Test to see if a valid channel string is converted into a Channel struct
	validChannel, err := NewChannel(validChannelPropertyString)
	if err != nil {
		t.Errorf("NewChannel(\"%v\"): Errored out with %v", validChannelPropertyString, err)
	} else {
		if validChannel.Cid != 1 || validChannel.Pid != 2 || validChannel.Name != "Sample Channel Name" || validChannel.TotalClients != 4 || validChannel.NeededSubscribePower != 5 {
			t.Errorf("NewChannel(\"%v\"): Parsed version %v does not match source input", validChannelPropertyString, validChannel)
		}
	}

	// Test to see if a valid channel string is converted into a Channel struct, (with null values)
	validNullParamChannel, err := NewChannel(validChannelPropertyStringWithNull)
	if err != nil {
		t.Errorf("NewChannel(\"%v\"): Errored out with %v", validChannelPropertyStringWithNull, err)
	} else {
		if validNullParamChannel.Cid != 1 || validNullParamChannel.Pid != 0 || validNullParamChannel.Name != "Sample Channel Name" || validNullParamChannel.TotalClients != 4 || validNullParamChannel.NeededSubscribePower != 5 {
			t.Errorf("NewChannel(\"%v\"): Parsed version %v does not match source input", validChannelPropertyStringWithNull, validNullParamChannel)
		}
	}

	// Test to see if a invalid channel string throws an error
	invalidChannel, err := NewChannel(invalidChannelPropertyString)
	if err == nil {
		t.Errorf("NewChannel(\"%v\"): Should have thrown an error. Instead received %v", invalidChannelPropertyString, invalidChannel)
	}

	// Test to make sure channel without cid is marked as new
	_, err = NewChannel(newChannelPropertyString)
	if err != nil {
		t.Errorf("NewChannel(\"%v\"): Errored out with %v", newChannelPropertyString, err)
	}
}

const validChannelUpdateString = "cid=2 pid=3"
const invalidChannelUpdateString = "cid=4 invalid=true"

func TestDeserialize(t *testing.T) {
	// Test to see if a valid channel string is converted into a Channel struct
	validChannel, err := NewChannel(validChannelPropertyString)
	if err != nil {
		t.Errorf("NewChannel(\"%v\"): Errored out with %v", validChannelPropertyString, err)
	}

	// Test a valid deserialize call
	_, err = validChannel.Deserialize(validChannelUpdateString)

	if err != nil {
		t.Errorf("channel.Deserialize(\"%v\"): Errored out with %v", validChannelUpdateString, err)
	}

	// Test an invalid deserialize call
	_, err = validChannel.Deserialize(invalidChannelUpdateString)

	if err == nil {
		t.Errorf("channel.Deserialize(\"%v\"): should have thrown an error", invalidChannelUpdateString, err)
	}
}

func TestSerialize(t *testing.T) {
	// Setup our test channel
	channel := &Channel{}

	// Cid should never be returned
	channel.Cid = 1
	propertyString, err := channel.Serialize()
	if err != nil {
		t.Errorf("channel.Serialize(%v): Errored out with %v", channel, err)
	}

	if strings.Index(propertyString, "cid=") != -1 {
		t.Errorf("channel.Serialize(%v): Included cid in returned property string(%v)", channel, propertyString)
	}

	// Serialized channel should provide all the data from the channel
	propertyString, err = channel.Serialize()
	if err != nil {
		t.Errorf("channel.Serialize(%v): Errored out with %v", channel, err)
	}

	// Pid
	channel.Pid = 2
	propertyString, err = channel.Serialize()
	if strings.Index(propertyString, fmt.Sprintf("pid=%d", channel.Pid)) == -1 {
		t.Errorf("channel.Serialize(%v): pid missing from returned property string \"%v\"", channel, propertyString)
	}

	// Order
	channel.Order = 3
	propertyString, err = channel.Serialize()
	if strings.Index(propertyString, fmt.Sprintf("channel_order=%d", channel.Order)) == -1 {
		t.Errorf("channel.Serialize(%v): channel_order missing from returned property string \"%v\"", channel, propertyString)
	}

	// Name
	channel.Name = "foo bar"
	propertyString, err = channel.Serialize()
	if strings.Index(propertyString, fmt.Sprintf("channel_name=%v", Escape(channel.Name))) == -1 {
		t.Errorf("channel.Serialize(%v): channel_name missing from returned property string \"%v\"", channel, propertyString)
	}

	// NeededSubscribePower
}
