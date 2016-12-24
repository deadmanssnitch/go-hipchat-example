# Dead Man's Snitch - HipChat Alerter

This is an example of using Dead Man's Snitch's webhooks integration to send
alerts to HipChat.

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Config

### HIPCHAT_TOKEN

Create a new [User API Token](https://hipchat.com/account/api) with the
"Send Notification" scope.

### HIPCHAT_ROOM

Grab the numeric ID of the room you want to post to from the room's URL. 

For example: use `12345` if your room's URL is `/chat/room/12345`.

### HIPCHAT_HOSTNAME (optional)

Ignore this unless you're an internally hosted HipChat Server instead of
hipchat.com.
