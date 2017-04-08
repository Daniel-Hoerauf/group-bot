package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"math/rand"
	"time"
)

var (
	secrets TokenFile
	Debug   *log.Logger
	Error   *log.Logger
)

func getBotId(group string) string {
	// Find the bot id associated with the group message recieved from
	for _, bot := range secrets.Bots {
		if bot.Group == group {
			return bot.BotId
		}
	}
	Error.Printf("No bot found for GroupID %s\n", group)
	return ""
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error.Printf("Could not read http body %s\n", err)
		return
	}
	message := body[:]
	response := GroupmeContent{}
	json.Unmarshal(message, &response)
	go giphy(response)
	go supreme(response)
	response_text, _ := json.MarshalIndent(response, "", "\t")
	Debug.Printf("Received message!\n%s", response_text)
}

func callGiphy(keywords []string, bot_id string) error {
	gif, err := downloadGif(keywords)
	if err != nil {
		Error.Printf("Could not download GIF: %s", err)
		return err
	}
	groupmeUrl, err := groupmeImageHost(gif)
	if err != nil {
		Error.Printf("Could not upload GIF: %s", err)
		return err
	}
	err = postGif(groupmeUrl+".large", bot_id)
	return err

}

func giphy(message GroupmeContent) {
	if message.Text == "" {
		return
	}
	tokens := strings.Split(message.Text, " ")
	if strings.ToLower(tokens[0]) == "/giphy" {
		for _, blacklisted := range secrets.BlackList {
			if blacklisted == message.SenderId {
				Debug.Printf("User blocked: %s\n", message.Name)
				return
			}
		}
		bot_id := getBotId(message.Group)
		escaped_tokens := make([]string, len(tokens)-1)
		for i, word := range tokens[1:] {
			escaped_tokens[i] = url.QueryEscape(word)
		}
		callGiphy(escaped_tokens, bot_id)

	}
}

func supreme(message GroupmeContent) {
	if (message.Sender == "bot") {
		return
	}
	var sup SUPREMEBot
	group := message.Group
	for _, bot := range secrets.SUPREME {
		if bot.Group == group {
			sup = bot
			break
		}
	}
	if (sup == SUPREMEBot{}) {
		return
	}
	val := rand.Intn(100)
	if (sup.Odds > val) {
		Debug.Printf("%+v\n", sup)
		postURL := "https://api.groupme.com/v3/bots/post"
		params := GroupmePost{sup.BotId, "SUPREME", []PostImg{}}
		binData, _ := json.Marshal(params)
		req, err := http.NewRequest("Post", postURL, bytes.NewBuffer(binData))
		if err != nil {
			Error.Printf("SUPREME error %s", err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			Error.Printf("SUPREME error %s", err)
		}
		defer resp.Body.Close()
	}


}

func postGif(imgLoc string, token string) error {
	postURL := "https://api.groupme.com/v3/bots/post"
	params := GroupmePost{token, "", []PostImg{PostImg{"image", imgLoc}}}
	binData, _ := json.Marshal(params)
	req, err := http.NewRequest("Post", postURL, bytes.NewBuffer(binData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	return err
}

func main() {
	Debug = log.New(os.Stdout, "GROUP_BOT: ", log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
	rand.Seed(time.Now().UTC().UnixNano())

	content, err := ioutil.ReadFile("secrets.json")
	if err != nil {
		Error.Printf("Cannot Read File: %s\n", err)
		os.Exit(1)
	}
	err = json.Unmarshal(content, &secrets)
	if err != nil {
		Error.Printf("JSON Parse Error: %s\n", err)
		os.Exit(1)
	}
	secrets_text, _ := json.MarshalIndent(secrets, "", "\t")
	Debug.Printf("Contents of secrets.json:\n%s", secrets_text)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
	Debug.Printf("Listening on port 80")
}
