package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var secrets TokenFile

func getBotId(group string) string {
	for _, bot := range secrets.Bots {
		if bot.Group == group {
			return bot.BotId
		}
	}
	return ""
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
	r.ParseForm()
	fmt.Println("Content received!")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	message := body[:]
	response := GroupmeContent{}
	json.Unmarshal(message, &response)
	go giphy(response)
	fmt.Println(response.Text)
}

func callGiphy(keywords []string, bot_id string) error {
	// file := strings.Join(keywords[:],"+") + ".gif"
	gif, err := downloadGif(keywords)
	if err != nil {
		fmt.Println("ERROR", err)
		return err
	}
	groupmeUrl, err := groupmeImageHost(gif)
	// defer destroyFile(file)
	if err != nil {
		return err
	}
	err = postGif(groupmeUrl+".large", bot_id)
	return err

}

func giphy(message GroupmeContent) {
	if message.Text == "" {
		return
	}
	for _, blacklisted := range secrets.BlackList {
		if blacklisted == message.SenderId {
			fmt.Printf("User blocked: %s\n", message.Name)
			return
		}
	}
	tokens := strings.Split(message.Text, " ")
	if tokens[0] == "/giphy" {
		bot_id := getBotId(message.Group)
		escaped_tokens := make([]string, len(tokens)-1)
		for i, word := range tokens[1:] {
			escaped_tokens[i] = url.QueryEscape(word)
		}
		fmt.Println(escaped_tokens)
		callGiphy(escaped_tokens, bot_id)

	}
}

func postGif(imgLoc string, token string) error {
	postURL := "https://api.groupme.com/v3/bots/post"
	params := GroupmePost{token, "", []PostImg{PostImg{"image", imgLoc}}}
	binData, _ := json.Marshal(params)
	fmt.Println(string(binData))
	req, err := http.NewRequest("Post", postURL, bytes.NewBuffer(binData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	return err
}

func main() {
	content, err := ioutil.ReadFile("secrets.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(content, &secrets)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(secrets.Token)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
