- A simple CLI tool to enable slack integration,
- Can be used for any CI/CD procedures, or any other automated pipelines.
- Simply download the binary and use it.

---

For any suggestions or issues, you are welcome to submit it [here](https://github.com/berkil/slack-terminal/issues/new)

---

[![Go Report Card](https://goreportcard.com/badge/github.com/berkil/slack-terminal)](https://goreportcard.com/report/github.com/berkil/slack-terminal)  

---

Current version: [0.2.0](https://github.com/berkil/slack-push-message/releases/tag/0.2.0)  
Download links:  
  * [OSX](https://github.com/berkil/slack-terminal/files/1650143/slack-cli_OSX.zip)
  * [Linux 64Bit Arch](https://github.com/berkil/slack-terminal/files/1650142/slack-cli_Linux_64.zip)
  * [Linux 32Bit Arch](https://github.com/berkil/slack-terminal/files/1650141/slack-cli_Linux_32.zip)  

---

Help and usage:
~~~
NAME:
   Slack Push Message - Enjoy sending messages through the terminal

USAGE:
   slack-cli [global options] [arguments...]

VERSION:
   0.2.0

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

Example:
  ~~~
    ./slack --channel "test_channel" \
      --webhook "https://hooks.slack.com/services/JUSTEXAMPLE/WEBHOOK/ABcdER5" \
      --username "test_bot" \
      --message "Test message \n New line with <https://github.com/berkil|url>" \
      --emoji ":ant:"
  ~~~
