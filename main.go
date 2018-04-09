package main

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/nlopes/slack"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	channelFlag   = kingpin.Flag("channel", "Slack Channel Name or ID").Short('c').Strings()
	tokenFlag     = kingpin.Flag("token", "Slack token").Short('t').String()
	titleFlag     = kingpin.Flag("title", "Message title").String()
	footerFlag    = kingpin.Flag("footer", "Footer to use").String()
	colorFlag     = kingpin.Flag("color", "Message color").String()
	usernameFlag  = kingpin.Flag("username", "Username to use").String()
	iconEmojiFlag = kingpin.Flag("icon_emoji", "Emoji to use as the icon").String()
	iconURLFlag   = kingpin.Flag("icon_url", "URL to an image to use as the icon").URL()
	teeModeFlag   = kingpin.Flag("tee", "Tee mode (pipe stdin to stdout)").Bool()
	messageArg    = kingpin.Arg("message", "Text of the message to send").String()
)

func main() {
	kingpin.Version("1.2.1")
	kingpin.Parse()

	if channelFlag == nil || len(*channelFlag) <= 0 {
		if env := os.Getenv("SLACK_CHANNEL"); len(env) > 0 {
			channelFlag = new([]string)
			parts := strings.Split(env, ",")
			for i := 0; i < len(parts); i++ {
				if part := strings.TrimSpace(parts[i]); len(part) > 0 {
					*channelFlag = append(*channelFlag, part)
				}
			}
		} else {
			fmt.Println("No environment variable `SLACK_CHANNEL' present")
			os.Exit(1)
		}
	}

	if tokenFlag == nil || len(*tokenFlag) <= 0 {
		if env := os.Getenv("SLACK_TOKEN"); len(env) > 0 {
			tokenFlag = new(string)
			*tokenFlag = env
		} else {
			fmt.Println("No environment variable `SLACK_TOKEN' present")
			os.Exit(1)
		}
	}
	*tokenFlag = strings.TrimSpace(*tokenFlag)

	if titleFlag == nil || len(*titleFlag) <= 0 {
		titleFlag = new(string)
		if env := os.Getenv("SLACK_TITLE"); len(env) > 0 {
			*titleFlag = env
		} else {
			*titleFlag = ""
		}
	}
	*titleFlag = strings.TrimSpace(*titleFlag)

	if footerFlag == nil || len(*footerFlag) <= 0 {
		footerFlag = new(string)
		if env := os.Getenv("SLACK_FOOTER"); len(env) > 0 {
			*footerFlag = env
		} else {
			*footerFlag = ""
		}
	}
	*footerFlag = strings.TrimSpace(*footerFlag)

	if colorFlag == nil || len(*colorFlag) <= 0 {
		colorFlag = new(string)
		if env := os.Getenv("SLACK_COLOR"); len(env) > 0 {
			*colorFlag = env
		} else {
			*colorFlag = ""
		}
	}
	*colorFlag = strings.TrimSpace(*colorFlag)

	if usernameFlag == nil || len(*usernameFlag) <= 0 {
		usernameFlag = new(string)
		if env := os.Getenv("SLACK_USERNAME"); len(env) > 0 {
			*usernameFlag = env
		} else {
			*usernameFlag = ""
		}
	}
	*usernameFlag = strings.TrimSpace(*usernameFlag)

	if iconEmojiFlag == nil || len(*iconEmojiFlag) <= 0 {
		iconEmojiFlag = new(string)
		if env := os.Getenv("SLACK_ICON_EMOJI"); len(env) > 0 {
			*iconEmojiFlag = env
		} else {
			*iconEmojiFlag = ""
		}
	}
	*iconEmojiFlag = strings.TrimSpace(*iconEmojiFlag)

	if iconURLFlag == nil || *iconURLFlag == nil || len((*iconURLFlag).String()) <= 0 {
		if env := os.Getenv("SLACK_ICON_URL"); len(env) > 0 {
			parsedUrl, err := url.Parse(strings.TrimSpace(env))
			if err == nil && parsedUrl != nil {
				iconURLFlag = &parsedUrl
			}
		}
	}

	// read message
	if messageArg == nil || len(*messageArg) <= 0 {
		var builder strings.Builder
		var err error
		if *teeModeFlag == true {
			_, err = io.Copy(io.MultiWriter(&builder, os.Stdout), os.Stdin)
		} else {
			_, err = io.Copy(&builder, os.Stdin)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		messageArg = new(string)
		*messageArg = builder.String()
	}

	if len(strings.TrimSpace(*messageArg)) <= 0 {
		return
	}

	// get channel ID
	var channelIDs []string
	api := slack.New(*tokenFlag)
	channels, err := api.GetChannels(true)
	if err != nil {
		fmt.Println("Unable to receive channel list:")
		fmt.Println(err)
		os.Exit(1)
	}

	for _, channel := range channels {
		for i := 0; i < len(*channelFlag); i++ {
			if strings.EqualFold(channel.NameNormalized, (*channelFlag)[i]) || strings.EqualFold(channel.Name, (*channelFlag)[i]) || strings.EqualFold(channel.ID, (*channelFlag)[i]) {
				channelIDs = append(channelIDs, channel.ID)
			}
		}
	}

	groups, err := api.GetGroups(true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to receive (private) channel list:\n")
		fmt.Fprintln(os.Stderr, err)
	} else {
		for _, channel := range groups {
			for i := 0; i < len(*channelFlag); i++ {
				if strings.EqualFold(channel.NameNormalized, (*channelFlag)[i]) || strings.EqualFold(channel.Name, (*channelFlag)[i]) || strings.EqualFold(channel.ID, (*channelFlag)[i]) {
					channelIDs = append(channelIDs, channel.ID)
				}
			}
		}
	}

	if len(channelIDs) <= 0 {
		fmt.Println("No channels found")
		os.Exit(1)
	}

	var text string

	parameters := slack.PostMessageParameters{
		AsUser: false,
	}

	if len(*titleFlag) > 0 || len(*colorFlag) > 0 || len(*footerFlag) > 0 || len(*messageArg) > 1024 {
		var attachment slack.Attachment

		if len(*titleFlag) > 0 {
			attachment.Title = *titleFlag
		}
		if len(*footerFlag) > 0 {
			attachment.Footer = *footerFlag
		}
		if len(*colorFlag) > 0 {
			attachment.Color = *colorFlag
		}
		attachment.Text = *messageArg

		parameters.Attachments = []slack.Attachment{
			attachment,
		}
	} else {
		text = *messageArg
	}

	if len(*usernameFlag) > 0 {
		parameters.Username = *usernameFlag
	}

	if len(*iconEmojiFlag) > 0 {
		parameters.IconEmoji = *iconEmojiFlag
	}

	if iconURLFlag != nil && *iconURLFlag != nil && len((*iconURLFlag).String()) > 0 {
		parameters.IconURL = (*iconURLFlag).String()
	}

	for i := 0; i < len(channelIDs); i++ {
		_, _, err := api.PostMessage(channelIDs[i], text, parameters)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
