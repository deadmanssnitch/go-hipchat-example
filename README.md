# Dead Man's Snitch - HipChat Alerter

This is an example of using Dead Man's Snitch's webhooks integration to send
alerts to HipChat.

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Config

### HIPCHAT_TOKEN (required)

Create a new [User API Token](https://hipchat.com/account/api) with the
"Send Notification" scope.

### HIPCHAT_ROOM (required)

Grab the numeric ID of the room you want to post to from the room's URL. 

For example: use `12345` if your room's URL is `/chat/room/12345`.

### HIPCHAT_HOSTNAME (optional)

Ignore this unless you're an internally hosted HipChat Server instead of
hipchat.com.

### DMS_PASSWORD (recommended)

Set a password (username actually) that will need to be sent via Basic auth to
authenticate Dead Man's Snitch. We recommend this as it adds an extra layer to
of protection for this service.

When adding the Webhook integration on deadmanssnitch.com, you add it as the
username in the URL. Example: https://password@some-app.herokuapp.com.
