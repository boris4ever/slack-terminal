---

Simple code written in go, to push messages to slack channel.

---

Current version: [0.1.2](https://github.com/berkil/slack-push-message/releases/tag/0.1.1)  
Download links:  
- First Parser Option:
  * [OSX](https://github.com/berkil/slack-terminal/files/1628708/slack_OSX.zip)  
  * [Linux 64Bit Arch](https://github.com/berkil/slack-terminal/files/1628707/slack_Linux_64.zip)  
  * [Linux 32Bit Arch](https://github.com/berkil/slack-terminal/files/1628706/slack_Linux_32.zip)  
Help text:
~~~
NAME:
   Slack Push Message

USAGE:
   ./slack -m <MESSAGE> -w <WEBHOOK> -u <USERNAME> -c <CHANNEL>

VERSION:
   0.1.2

AUTHOR:
   Boris Bakshiyev

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --message Message, -m Message     Message to be sent
   --webhook webhook, -w webhook     The slack webhook
   --username username, -u username  The username which will be shown as the sender
   --channel channel, -c channel     The channel to send the message too
   --help, -h                        show help
   --version, -v                     print the version
~~~


- Second Parser Option:
  * [OSX](https://github.com/berkil/slack-terminal/files/1628712/slack_OSX.zip)
  * [Linux 64Bit Arch](https://github.com/berkil/slack-terminal/files/1628711/slack_Linux_64.zip)  
  * [Linux 32Bit Arch](https://github.com/berkil/slack-terminal/files/1628710/slack_Linux_32.zip)  
Help Text:
~~~
usage: slack --message=MESSAGE --webhook=WEBHOOK --username=USERNAME --channel=CHANNEL [<flags>]

Flags:
  -h, --help               Show context-sensitive help (also try --help-long and --help-man).
  -m, --message=MESSAGE    Message to be sent
  -w, --webhook=WEBHOOK    The slack webhook
  -u, --username=USERNAME  The username which will be shown as the sender
  -c, --channel=CHANNEL    The channel to send the message too
~~~

---

Parameters:  
  * **channel**: slack channel.  
  * **username**: Which username will be shown for the message.  
  * **webhook**: The slack [incoming webhook](https://api.slack.com/incoming-webhooks).  
  * **message**: The message you want to send.  

---

Usage:
  `./slack --channel="<CHANNEL>" --webhook="<WEB_HOOK>" --username="<USER_NAME>" --message="<MESSAGE>"`

Example:
  ~~~
    ./slack --channel="test_channel" \
      --webhook="https://hooks.slack.com/services/JUSTEXAMPLE/WEBHOOK/ABcdER5" \
      --username="test_bot" \
      --message="Test message \n New line"
  ~~~
