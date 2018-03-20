# 2slack

Send a message to slack

```bash
echo "Hello World" | 2slack
```

## Usage
```bash
usage: 2slack [<flags>] [<message>]

Flags:
      --help                 Show context-sensitive help (also try --help-long and --help-man).
  -c, --channel=CHANNEL ...  Slack Channel Name or ID
  -t, --token=TOKEN          Slack token
      --title=TITLE          Message title
      --footer=FOOTER        Footer to use
      --color=COLOR          Message color
      --username=USERNAME    Username to use
      --version              Show application version.

Args:
  [<message>]  Message text
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

export SLACK_CHANNEL = Channel1
export SLACK_TOKEN = MySlackToken
export SLACK_TITLE = Date
export SLACK_COLOR = green

date | 2slack
2slack "Hello World!"
```
