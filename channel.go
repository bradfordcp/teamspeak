package teamspeak

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Channel struct {
	// Retrieved in ChannelList call
	Cid                         uint
	Pid                         uint
	ChannelOrder                uint
	ChannelName                 string
	TotalClients                uint
	ChannelNeededSubscribePower uint

	// Retrieved in ChannelInfo call
	ChannelTopic                         string
	ChannelDescription                   string
	ChannelPassword                      string
	ChannelCodec                         uint
	ChannelCodecQuality                  uint
	ChannelMaxClients                    int
	ChannelMaxFamilyClients              int
	ChannelFlagPermanent                 bool
	ChannelFlagSemiPermanent             bool
	ChannelFlagDefault                   bool
	ChannelFlagPassword                  bool
	ChannelCodecLatencyFactor            uint
	ChannelCodecIsUnencrypted            bool
	ChannelSecuritySalt                  string
	ChannelDeleteDelay                   uint
	ChannelFlagMaxClientsUnlimited       bool
	ChannelFlagMaxFamilyClientsUnlimited bool
	ChannelFlagMaxFamilyClientsInherited bool
	ChannelFilepath                      string
	ChannelNeededTalkPower               uint
	ChannelForcedSilence                 bool
	ChannelNamePhonetic                  string
	ChannelIconId                        int
	ChannelFlagPrivate                   bool
	SecondsEmpty                         int
}

func NewChannel(channelStr string) (*Channel, error) {
	// Instantiate the Channel
	channel := &Channel{}

	// Update the properties with the passed in data
	channel.Deserialize(channelStr)

	// Sanity check on the channel
	if channel.Cid == 0 {
		return channel, errors.New(fmt.Sprintf("Invalid channel, no cid present from %v", channelStr))
	}

	return channel, nil
}

// Update the properties of the channel with the attributes passed in
func (channel *Channel) Deserialize(propertiesStr string) (*Channel, error) {
	// Split the tokens and fill in our channel
	tokens := strings.Split(propertiesStr, " ")
	for _, token := range tokens {
		attribute := strings.SplitN(token, "=", 2)

		if len(attribute) == 2 {
			switch attribute[0] {
			case "cid":
				cid, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.Cid = uint(cid)

			case "pid":
				pid, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.Pid = uint(pid)

			case "channel_order":
				channelOrder, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.ChannelOrder = uint(channelOrder)

			case "channel_name":
				channel.ChannelName = attribute[1]

			case "total_clients":
				totalClients, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.TotalClients = uint(totalClients)

			case "channel_needed_subscribe_power":
				channelNeededSubscribePower, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.ChannelNeededSubscribePower = uint(channelNeededSubscribePower)

			case "channel_topic":
				channel.ChannelTopic = attribute[1]

			case "channel_description":
				channel.ChannelDescription = attribute[1]

			case "channel_password":
				channel.ChannelPassword = attribute[1]

			case "channel_codec":
				channelCodec, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.ChannelCodec = uint(channelCodec)

			case "channel_codec_quality":
				channelCodecQuality, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.ChannelCodecQuality = uint(channelCodecQuality)

			case "channel_maxclients":
				channelMaxClients, err := strconv.ParseInt(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.ChannelMaxClients = int(channelMaxClients)

			case "channel_maxfamilyclients":
				channelMaxFamilyClients, err := strconv.ParseInt(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.ChannelMaxFamilyClients = int(channelMaxFamilyClients)

			case "channel_flag_permanent":
				channelFlagPermanent, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.ChannelFlagPermanent = bool(channelFlagPermanent)

			case "channel_flag_semi_permanent":
				channelFlagSemiPermanent, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.ChannelFlagSemiPermanent = bool(channelFlagSemiPermanent)

			case "channel_flag_default":
				channelFlagDefault, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.ChannelFlagDefault = bool(channelFlagDefault)

			case "channel_flag_password":
				channelFlagPassword, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.ChannelFlagPassword = bool(channelFlagPassword)

			case "channel_codec_latency_factor":
				channelCodecLatencyFactor, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.ChannelCodecLatencyFactor = uint(channelCodecLatencyFactor)

			case "channel_codec_is_unencrypted":
				channelCodecIsUnencrypted, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.ChannelCodecIsUnencrypted = bool(channelCodecIsUnencrypted)

			case "channel_security_salt":
				channel.ChannelSecuritySalt = attribute[1]

			case "channel_delete_delay":
				channelDeleteDelay, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.ChannelDeleteDelay = uint(channelDeleteDelay)

			case "channel_flag_maxclients_unlimited":
				channelFlagMaxClientsUnlimited, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.ChannelFlagMaxClientsUnlimited = bool(channelFlagMaxClientsUnlimited)

			case "channel_flag_maxfamilyclients_unlimited":
				channelFlagMaxFamilyClientsUnlimited, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.ChannelFlagMaxFamilyClientsUnlimited = bool(channelFlagMaxFamilyClientsUnlimited)

			case "channel_flag_maxfamilyclients_inherited":
				channelFlagMaxFamilyClientsInherited, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.ChannelFlagMaxFamilyClientsInherited = bool(channelFlagMaxFamilyClientsInherited)

			case "channel_filepath":
				channel.ChannelFilepath = attribute[1]

			case "channel_needed_talk_power":
				channelNeededTalkPower, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.ChannelNeededTalkPower = uint(channelNeededTalkPower)

			case "channel_forced_silence":
				channelForcedSilence, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.ChannelForcedSilence = bool(channelForcedSilence)

			case "channel_name_phonetic":
				channel.ChannelNamePhonetic = attribute[1]

			case "channel_icon_id":
				channelIconId, err := strconv.ParseInt(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.ChannelIconId = int(channelIconId)

			case "channel_flag_private":
				channelFlagPrivate, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.ChannelFlagPrivate = bool(channelFlagPrivate)

			case "seconds_empty":
				secondsEmpty, err := strconv.ParseInt(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.SecondsEmpty = int(secondsEmpty)

			default:
				return channel, errors.New(fmt.Sprintf("Error invalid parameter detected (%v) from %v", attribute[0], propertiesStr))
			}
		}
	}

	return channel, nil
}

// Reads the list of channels
func (ts3 *Conn) ChannelList() ([]*Channel, error) {
	response, err := ts3.SendCommand("channellist")
	if ts3Err, ok := err.(*Error); ok && ts3Err.Id == 0 {
		// Split the channel data on the | character
		rawChannels := strings.Split(response, "|")
		channels := make([]*Channel, len(rawChannels))

		// Iterate over each raw channel and parse the attributes
		for i, rawChannel := range rawChannels {
			channel, err := NewChannel(rawChannel)
			if err != nil {
				return channels, err
			}
			channels[i] = channel
		}

		return channels, nil
	}

	empty := make([]*Channel, 0)
	return empty, err
}

func (ts3 *Conn) ChannelInfo(channel *Channel) error {
	response, err := ts3.SendCommand(fmt.Sprintf("channelinfo cid=%d", channel.Cid))
	if ts3Err, ok := err.(*Error); ok && ts3Err.Id == 0 {
		_, err := channel.Deserialize(response)

		if err != nil {
			return err
		}
	}

	return nil
}
