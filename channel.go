package teamspeak

import (
	"strconv"
	"strings"
)

type Channel struct {
	Cid                         uint
	Pid                         uint
	ChannelOrder                uint
	ChannelName                 string
	TotalClients                uint
	ChannelNeededSubscribePower uint
}

func NewChannel(channelStr string) (*Channel, error) {
	channel := &Channel{}

	tokens := strings.Split(channelStr, " ")
	for _, token := range tokens {
		attribute := strings.Split(token, "=")
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
