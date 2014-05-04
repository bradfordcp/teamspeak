package teamspeak

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Channel struct {
	// Retrieved in ChannelList call
	Cid                  uint
	Pid                  uint
	Order                uint
	Name                 string
	TotalClients         uint
	NeededSubscribePower uint

	// Retrieved in ChannelInfo call
	Topic                         string
	Description                   string
	Password                      string
	Codec                         uint
	CodecQuality                  uint
	MaxClients                    int
	MaxFamilyClients              int
	FlagPermanent                 bool
	FlagSemiPermanent             bool
	FlagDefault                   bool
	FlagPassword                  bool
	CodecLatencyFactor            uint
	CodecIsUnencrypted            bool
	SecuritySalt                  string
	DeleteDelay                   uint
	FlagMaxClientsUnlimited       bool
	FlagMaxFamilyClientsUnlimited bool
	FlagMaxFamilyClientsInherited bool
	Filepath                      string
	NeededTalkPower               uint
	ForcedSilence                 bool
	NamePhonetic                  string
	IconId                        int
	FlagPrivate                   bool
	SecondsEmpty                  int

	// Tracking state for Save calls
	newRecord  bool
	infoLoaded bool
}

func NewChannel(channelStr string) (*Channel, error) {
	// Instantiate the Channel
	channel := &Channel{}

	// Update the properties with the passed in data
	_, err := channel.Deserialize(channelStr)
	if err != nil {
		return channel, err
	}

	// Sanity check on the channel
	if channel.Cid == 0 {
		channel.newRecord = true
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
				order, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.Order = uint(order)

			case "channel_name":
				channel.Name = Unescape(attribute[1])

			case "total_clients":
				totalClients, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.TotalClients = uint(totalClients)

			case "channel_needed_subscribe_power":
				neededSubscribePower, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.NeededSubscribePower = uint(neededSubscribePower)

			case "channel_topic":
				channel.Topic = Unescape(attribute[1])

			case "channel_description":
				channel.Description = Unescape(attribute[1])

			case "channel_password":
				channel.Password = Unescape(attribute[1])

			case "channel_codec":
				codec, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.Codec = uint(codec)

			case "channel_codec_quality":
				codecQuality, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.CodecQuality = uint(codecQuality)

			case "channel_maxclients":
				maxClients, err := strconv.ParseInt(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.MaxClients = int(maxClients)

			case "channel_maxfamilyclients":
				maxFamilyClients, err := strconv.ParseInt(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.MaxFamilyClients = int(maxFamilyClients)

			case "channel_flag_permanent":
				flagPermanent, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.FlagPermanent = bool(flagPermanent)

			case "channel_flag_semi_permanent":
				flagSemiPermanent, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.FlagSemiPermanent = bool(flagSemiPermanent)

			case "channel_flag_default":
				flagDefault, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.FlagDefault = bool(flagDefault)

			case "channel_flag_password":
				flagPassword, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.FlagPassword = bool(flagPassword)

			case "channel_codec_latency_factor":
				codecLatencyFactor, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.CodecLatencyFactor = uint(codecLatencyFactor)

			case "channel_codec_is_unencrypted":
				codecIsUnencrypted, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.CodecIsUnencrypted = bool(codecIsUnencrypted)

			case "channel_security_salt":
				channel.SecuritySalt = Unescape(attribute[1])

			case "channel_delete_delay":
				deleteDelay, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.DeleteDelay = uint(deleteDelay)

			case "channel_flag_maxclients_unlimited":
				flagMaxClientsUnlimited, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.FlagMaxClientsUnlimited = bool(flagMaxClientsUnlimited)

			case "channel_flag_maxfamilyclients_unlimited":
				flagMaxFamilyClientsUnlimited, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.FlagMaxFamilyClientsUnlimited = bool(flagMaxFamilyClientsUnlimited)

			case "channel_flag_maxfamilyclients_inherited":
				flagMaxFamilyClientsInherited, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.FlagMaxFamilyClientsInherited = bool(flagMaxFamilyClientsInherited)

			case "channel_filepath":
				channel.Filepath = Unescape(attribute[1])

			case "channel_needed_talk_power":
				neededTalkPower, err := strconv.ParseUint(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.NeededTalkPower = uint(neededTalkPower)

			case "channel_forced_silence":
				forcedSilence, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.ForcedSilence = bool(forcedSilence)

			case "channel_name_phonetic":
				channel.NamePhonetic = Unescape(attribute[1])

			case "channel_icon_id":
				iconId, err := strconv.ParseInt(attribute[1], 10, 32)
				if err != nil {
					return channel, err
				}
				channel.IconId = int(iconId)

			case "channel_flag_private":
				flagPrivate, err := strconv.ParseBool(attribute[1])
				if err != nil {
					return channel, err
				}
				channel.FlagPrivate = bool(flagPrivate)

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

// Pull additional channel info
func (ts3 *Conn) ChannelInfo(channel *Channel) error {
	response, err := ts3.SendCommand(fmt.Sprintf("channelinfo cid=%d", channel.Cid))
	if ts3Err, ok := err.(*Error); ok && ts3Err.Id == 0 {
		_, err := channel.Deserialize(response)

		if err != nil {
			return err
		}

		// Set the flag telling us ChannelInfo was called on this channel
		channel.infoLoaded = true
	}

	return nil
}
