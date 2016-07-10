#Giphy Bot for Groupme
##Installation
To install this program, run:
```
go get -u github.com/Daniel-Hoerauf/group-bot
```
##Usage
###Hosting
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
    ]
}
```
You will need to add a new entry under bots for each group you wish to have a bot listening.
###Using
To use the bot once it is running and listening any post in the groups of the format:
```
/giphy any random phrase
```
will call the [giphy](giphy.com) api and post the resulting gif back into the group
