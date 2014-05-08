package teamspeak

import (
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
const validChannelInfoString = "pid=2 channel_name=Foo\\sBar\\sBaz channel_topic channel_description=Multi-line\\nDescription channel_password=apassword channel_codec=3 channel_codec_quality=4 channel_maxclients=-1 channel_maxfamilyclients=-1 channel_order=5 channel_flag_permanent=1 channel_flag_semi_permanent=0 channel_flag_default=0 channel_flag_password=0 channel_codec_latency_factor=1 channel_codec_is_unencrypted=1 channel_security_salt channel_delete_delay=0 channel_flag_maxclients_unlimited=1 channel_flag_maxfamilyclients_unlimited=0 channel_flag_maxfamilyclients_inherited=1 channel_filepath=files\\/virtualserver_1\\/channel_1 channel_needed_talk_power=0 channel_forced_silence=0 channel_name_phonetic channel_icon_id=4 channel_flag_private=0 seconds_empty=5000"

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
	} else {
		// Validate the updated values
		if validChannel.Cid != 2 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update Cid value", validChannelUpdateString)
		}
		if validChannel.Pid != 3 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update Pid value", validChannelUpdateString)
		}
	}

	// Test an invalid deserialize call
	_, err = validChannel.Deserialize(invalidChannelUpdateString)

	if err == nil {
		t.Errorf("channel.Deserialize(\"%v\"): should have thrown an error", invalidChannelUpdateString, err)
	}

	// Test a larger ChannelInfo Deserialization
	_, err = validChannel.Deserialize(validChannelInfoString)
	if err != nil {
		t.Errorf("channel.Deserialize(\"%v\"): Errored out with %v", validChannelInfoString, err)
	} else {
		if validChannel.Pid != 2 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update Pid value", validChannelInfoString)
		}
		if validChannel.Name != "Foo Bar Baz" {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update Name value", validChannelInfoString)
		}
		if validChannel.Topic != "" {
			t.Errorf("channel.Deserialize(\"%v\"): Updated Topic value, when it should not have", validChannelInfoString)
		}
		if validChannel.Description != "Multi-line\nDescription" {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update Description value", validChannelInfoString)
		}
		if validChannel.Password != "apassword" {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update Password value", validChannelInfoString)
		}
		if validChannel.Codec != 3 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update Codec value", validChannelInfoString)
		}
		if validChannel.CodecQuality != 4 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update Codec value", validChannelInfoString)
		}
		if validChannel.MaxClients != -1 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update MaxClients value", validChannelInfoString)
		}
		if validChannel.MaxFamilyClients != -1 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update MaxFamilyClients value", validChannelInfoString)
		}
		if validChannel.Order != 5 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update Order value", validChannelInfoString)
		}
		if validChannel.FlagPermanent != true {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update FlagPermanent value", validChannelInfoString)
		}
		if validChannel.FlagSemiPermanent != false {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update FlagSemiPermanent value", validChannelInfoString)
		}
		if validChannel.FlagDefault != false {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update FlagDefault value", validChannelInfoString)
		}
		if validChannel.FlagPassword != false {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update FlagPassword value", validChannelInfoString)
		}
		if validChannel.CodecLatencyFactor != 1 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update CodecLatencyFactor value", validChannelInfoString)
		}
		if validChannel.CodecIsUnencrypted != true {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update CodecIsUnencrypted value", validChannelInfoString)
		}
		if validChannel.SecuritySalt != "" {
			t.Errorf("channel.Deserialize(\"%v\"): Updated SecuritySalt value, when it should not have", validChannelInfoString)
		}
		if validChannel.DeleteDelay != 0 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update DeleteDelay value", validChannelInfoString)
		}
		if validChannel.FlagMaxClientsUnlimited != true {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update FlagMaxClientsUnlimited value", validChannelInfoString)
		}
		if validChannel.FlagMaxFamilyClientsUnlimited != false {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update FlagMaxFamilyClientsUnlimited value", validChannelInfoString)
		}
		if validChannel.FlagMaxFamilyClientsInherited != true {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update FlagMaxFamilyClientsInherited value", validChannelInfoString)
		}
		if validChannel.Filepath != "files/virtualserver_1/channel_1" {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update FilePath value", validChannelInfoString)
		}
		if validChannel.NeededTalkPower != 0 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update NeededTalkPower value", validChannelInfoString)
		}
		if validChannel.ForcedSilence != false {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update ForcedSilence value", validChannelInfoString)
		}
		if validChannel.NamePhonetic != "" {
			t.Errorf("channel.Deserialize(\"%v\"): Updated NamePhonetic value, when it should not have", validChannelInfoString)
		}
		if validChannel.IconId != 4 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update IconId value", validChannelInfoString)
		}
		if validChannel.FlagPrivate != false {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update FlagPrivate value", validChannelInfoString)
		}
		if validChannel.SecondsEmpty != 5000 {
			t.Errorf("channel.Deserialize(\"%v\"): Did not update SecondsEmpty value", validChannelInfoString)
		}
	}
}

func TestSerialize(t *testing.T) {
	// Setup our test channel
	channel := &Channel{}
	channel.Cid = 1
	channel.Name = "foo bar baz"
	channel.FlagPermanent = true
	channel.MaxClients = -1

	// Test Uiont
	propertyString, err := channel.Serialize("Cid")
	if err != nil {
		t.Errorf("channel.Serialize(\"Cid\"): Errored out with %v", err)
	}

	if propertyString != "cid=1" {
		t.Errorf("channel.Serialize(\"Cid\"): Did not include cid in returned property string(%v)", propertyString)
	}

	// Test String
	propertyString, err = channel.Serialize("Name")
	if err != nil {
		t.Errorf("channel.Serialize(\"Name\"): Errored out with %v", err)
	}

	if propertyString != "channel_name=foo\\sbar\\sbaz" {
		t.Errorf("channel.Serialize(\"Name\"): Did not include channel_name in returned property string(%v)", propertyString)
	}

	// Test bool
	propertyString, err = channel.Serialize("FlagPermanent")
	if err != nil {
		t.Errorf("channel.Serialize(\"FlagPermanent\"): Errored out with %v", err)
	}

	if propertyString != "channel_flag_permanent=1" {
		t.Errorf("channel.Serialize(\"FlagPermanent\"): Did not include channel_flag_permanent in returned property string(%v)", propertyString)
	}

	// Test int
	propertyString, err = channel.Serialize("MaxClients")
	if err != nil {
		t.Errorf("channel.Serialize(\"MaxClients\"): Errored out with %v", err)
	}

	if propertyString != "channel_maxclients=-1" {
		t.Errorf("channel.Serialize(\"MaxClients\"): Did not include channel_max_clients in returned property string(%v)", propertyString)
	}
}
