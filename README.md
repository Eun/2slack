# 2slack

Send a message to slack

```bash
echo "Hello World" | 2slack
```

## Usage
```bash
usage: 2slack [<flags>] [<message>]

Flags:
      --help                   Show context-sensitive help (also try --help-long and --help-man).
  -c, --channel=CHANNEL ...    Slack Channel Name or ID
  -t, --token=TOKEN            Slack token
      --title=TITLE            Message title
      --footer=FOOTER          Footer to use
      --color=COLOR            Message color
      --username=USERNAME      Username to use
      --icon_emoji=ICON_EMOJI  Emoji to use as the icon
      --icon_url=ICON_URL      URL to an image to use as the icon
      --tee                    Tee mode (pipe stdin to stdout)
      --version                Show application version.

Args:
  [<message>]  Text of the message to send

```

# Installation
You find releases in [Github Releases](https://github.com/Eun/2slack/releases) section.

Or you can use `go get`:
```bash
go get github.com/Eun/2slack
```

You can find instructions for slack [here](slack/README.md).

# Examples

```bash
2slack --channel=Channel1 --token=MySlackToken --title=Title --color=green "Hello from 2slack"
date | 2slack --channel=Channel1 --token=MySlackToken --title=Title --color=ff0000
```

## Using environment variables
```bash

export SLACK_CHANNEL=Channel1
export SLACK_TOKEN=MySlackToken
export SLACK_TITLE=Date
export SLACK_COLOR=green
export SLACK_USERNAME=2slack

date | 2slack
2slack "Hello World!"
```

## List of environment variables
| Name               | Commandline Equivalent |
|--------------------|------------------------|
| `SLACK_CHANNEL`    | `channel`              |
| `SLACK_TOKEN`      | `token`                |
| `SLACK_TITLE`      | `title`                |
| `SLACK_FOOTER`     | `footer`               |
| `SLACK_COLOR`      | `color`                |
| `SLACK_USERNAME`   | `username`             |
| `SLACK_ICON_EMOJI` | `icon_emoji`           |
| `SLACK_ICON_URL`   | `icon_url`             |


# Changelog
1.2.1: Added `tee` mode
1.2.0: Added `icon_emoji` & `icon_url`

1.0.0: Initial release