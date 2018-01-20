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
   Slack Push Message - Enjoy sending messages through the terminal

USAGE:
   /tmp/go-build765099906/command-line-arguments/_obj/exe/main [global options] [arguments...]

VERSION:
   0.1.2

AUTHOR:
   Boris Bakshiyev

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --message message, -m message     The message to be sent
   --webhook webhook, -w webhook     The slack webhook
   --username username, -u username  The username which will be shown as the sender
   --channel channel, -c channel     The channel to send the message too
   --help, -h                        show help
   --version, -v                     print the version
~~~

---

Parameters:  
  * **message**:   The message you want to send.  
  * **webhook**:   The slack [incoming webhook](https://api.slack.com/incoming-webhooks).
  * **username**:  Which username will be shown for the message.    
  * **channel**:   slack channel.  

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
