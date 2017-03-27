# Giphy Bot for Groupme
## Installation
To install this program, run:
```
go get -u github.com/Daniel-Hoerauf/group-bot
```
## Usage
### Hosting
In order to use the bot in Groupme, you must first create a bot on the [Groupme Site](https://dev.groupme.com/bots). Set the callback url to the url of the server you will be hosting on.
Once you have your access token and the bot created, you must create a secrets.json file
```secrets.json
{
    "access_token": "Ady1BaSb1...",
    "bots": [
        {
            "group": "20745774",
            "bot_id": "435b83..."
        }
    ],
    "blacklist": [
      "14320..."
    ]
}
```
You will need to add a new entry under bots for each group you wish to have a bot listening.

#### Blacklist
If there are any users who you would like to block from being able to use the bot, you can do so by adding their UserId under `blacklist` at the top level of the json

### Using
To use the bot once it is running and listening any post in the groups of the format:
```
/giphy any random phrase
```
will call the [giphy](giphy.com) api and post the resulting gif back into the group
