---

Simple code written in go, to push messages to slack channel.

---

Current version: [0.1.1](https://github.com/berkil/slack-push-message/releases/tag/0.1.1)  
Download links:  
- First Parser Option:
  * [OSX](https://github.com/berkil/slack-push-message/files/1628365/slack_OSX.zip)
  * [Linux 64Bit Arch](https://github.com/berkil/slack-push-message/files/1628360/slack_Linux_64.zip)
  * [Linux 32Bit Arch](https://github.com/berkil/slack-push-message/files/1628355/slack_Linux_32.zip)



- Second Parser Option:
  * [OSX](https://github.com/berkil/slack-push-message/files/1628366/slack_OSX.zip)
  * [Linux 64Bit Arch](https://github.com/berkil/slack-push-message/files/1628362/slack_Linux_64.zip)
  * [Linux 32Bit Arch](https://github.com/berkil/slack-push-message/files/1628356/slack_Linux_32.zip)

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
