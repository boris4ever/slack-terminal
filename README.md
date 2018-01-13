---

Simple code written in go, to push messages to slack channel.

---

Current version: [0.1](https://github.com/berkil/slack-push-message/releases/tag/0.1)  
Download links:  
  * [OSX](https://github.com/berkil/slack-push-message/files/1528417/slack_OSX.zip)
  * [Linux 64Bit Arch](https://github.com/berkil/slack-push-message/files/1528422/slack_Linux_64.zip)
  * [Linux 32Bit Arch](https://github.com/berkil/slack-push-message/files/1528423/slack_Linux_32.zip)

---

Parameters:  
  * **channel**: slack channel.  
  * **username**: Which username will be shown for the message.  
  * **url**: The slack [incoming webhook](https://api.slack.com/incoming-webhooks).  
  * **msg**: The message you want to send.  
  
---

Usage:
  `./slack -channel="<CHANNEL>" -url="<WEB_HOOK>" -username="<USER_NAME>" -msg="<MESSAGE>"`

Example:
  ~~~
    ./slack -channel="test_channel" \
      -url="https://hooks.slack.com/services/JUSTEXAMPLE/WEBHOOK/ABcdER5" \
      -username="test_bot" \
      -msg="Test message \n New line"
  ~~~
