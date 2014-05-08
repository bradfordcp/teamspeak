package teamspeak

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Channel struct {
	// Retrieved in ChannelList call
	Cid                  uint   `sq:"cid"`
	Pid                  uint   `sq:"pid"`
	Order                uint   `sq:"channel_order"`
	Name                 string `sq:"channel_name"`
	TotalClients         uint   `sq:"total_clients"`
	NeededSubscribePower uint   `sq:"channel_needed_subscribe_power"`

	// Retrieved in ChannelInfo call
	Topic                         string `sq:"channel_topic"`
	Description                   string `sq:"channel_description"`
	Password                      string `sq:"channel_password"`
	Codec                         uint   `sq:"channel_codec"`
	CodecQuality                  uint   `sq:"channel_codec_quality"`
	MaxClients                    int    `sq:"channel_maxclients"`
	MaxFamilyClients              int    `sq:"channel_maxfamilyclients"`
	FlagPermanent                 bool   `sq:"channel_flag_permanent"`
	FlagSemiPermanent             bool   `sq:"channel_flag_semi_permanent"`
	FlagDefault                   bool   `sq:"channel_flag_default"`
	FlagPassword                  bool   `sq:"channel_flag_password"`
	CodecLatencyFactor            uint   `sq:"channel_codec_latency_factor"`
	CodecIsUnencrypted            bool   `sq:"channel_codec_is_unencrypted"`
	SecuritySalt                  string `sq:"channel_security_salt"`
	DeleteDelay                   uint   `sq:"channel_delete_delay"`
	FlagMaxClientsUnlimited       bool   `sq:"channel_flag_maxclients_unlimited"`
	FlagMaxFamilyClientsUnlimited bool   `sq:"channel_flag_maxfamilyclients_unlimited"`
	FlagMaxFamilyClientsInherited bool   `sq:"channel_flag_maxfamilyclients_inherited"`
	Filepath                      string `sq:"channel_filepath"`
	NeededTalkPower               uint   `sq:"channel_needed_talk_power"`
	ForcedSilence                 bool   `sq:"channel_forced_silence"`
	NamePhonetic                  string `sq:"channel_name_phonetic"`
	IconId                        int    `sq:"channel_icon_id"`
	FlagPrivate                   bool   `sq:"channel_flag_private"`
	SecondsEmpty                  int    `sq:"seconds_empty"`
}

func NewChannel(channelStr string) (*Channel, error) {
	// Instantiate the Channel
	channel := &Channel{}

	// Update the properties with the passed in data
	_, err := channel.Deserialize(channelStr)
	if err != nil {
		return channel, err
	}

	return channel, nil
}

// Update the properties of the channel with the attributes passed in
func (channel *Channel) Deserialize(propertiesStr string) (*Channel, error) {
	channelType := reflect.ValueOf(channel).Elem()

	// Split the tokens and fill in our channel
	tokens := strings.Split(propertiesStr, " ")
	for _, token := range tokens {
		attribute := strings.SplitN(token, "=", 2)

		if len(attribute) == 2 {
			fieldFound := false

			// Loop through the fields on Channel and assign the field
			for i := 0; i < channelType.NumField(); i++ {
				field := channelType.Field(i)
				fieldTag := channelType.Type().Field(i).Tag

				// See if the attribute matches the "sq" tag on the struct field
				if attribute[0] == fieldTag.Get("sq") {
					fieldFound = true

					// Base on the type of field parse appropriately
					switch field.Kind() {
					case reflect.Uint:
						value, err := strconv.ParseUint(attribute[1], 10, 32)
						if err != nil {
							return channel, err
						}

						field.SetUint(value)

					case reflect.Int:
						value, err := strconv.ParseInt(attribute[1], 10, 32)
						if err != nil {
							return channel, err
						}
						field.SetInt(value)

					case reflect.Bool:
						value, err := strconv.ParseBool(attribute[1])
						if err != nil {
							return channel, err
						}
						field.SetBool(value)

					case reflect.String:
						field.SetString(Unescape(attribute[1]))

					default:
						return channel, errors.New(fmt.Sprintf("Cannot handle valid parameter (%v) type %v not supported", attribute[0], field.Kind()))
					}

					break
				}
			}

			// If the field is not found, raise an error
			if !fieldFound {
				return channel, errors.New(fmt.Sprintf("Error invalid parameter detected (%v) from %v", attribute[0], propertiesStr))
			}
		}
	}

	return channel, nil
}

func (channel *Channel) Serialize() (string, error) {
	properties := make([]string, 25)

	properties[0] = fmt.Sprintf("pid=%d", channel.Pid)
	properties[1] = fmt.Sprintf("channel_order=%d", channel.Order)
	properties[2] = fmt.Sprintf("channel_name=%v", Escape(channel.Name))

	return strings.Join(properties, " "), nil
}

// Reads the list of channels
func (ts3 *Connection) ChannelList() ([]*Channel, error) {
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
func (ts3 *Connection) ChannelInfo(channel *Channel) error {
	response, err := ts3.SendCommand(fmt.Sprintf("channelinfo cid=%d", channel.Cid))
	if ts3Err, ok := err.(*Error); ok && ts3Err.Id == 0 {
		_, err := channel.Deserialize(response)

		if err != nil {
			return err
		}
	}

	return nil
}

// Saves the Channel, for now this will push up all stored attributes including ones that have not changed
func (ts3 *Connection) ChannelEdit(channel *Channel) error {
	// Serialize the channel's properties
	propertyString, err := channel.Serialize()
	if err != nil {
		return err
	}

	// Call the channel edit command
	ts3.SendCommand(fmt.Sprintf("channeledit cid=%d %v", channel.Cid, propertyString))

	return nil
}
