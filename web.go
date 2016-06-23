package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
	"bytes"
	"strings"

)

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
	go giphy(response.Text)
	fmt.Println(response.Text)
}

func giphy(message string) {
	if message == "" {
		return
	}
	tokens := strings.Split(message, " ")
	if tokens[0] == "/giphy" {
		escaped_tokens := make([]string, len(tokens) - 1)
		for i, word := range tokens[1:] {
			escaped_tokens[i] = url.QueryEscape(word)
		}
		fmt.Println(escaped_tokens)
		callGiphy(escaped_tokens)

	}
}

func postGif(imgLoc string) (error) {
	postURL := "https://api.groupme.com/v3/bots/post"
	token := os.Getenv("GROUPME_BOT_TOKEN")
	params := GroupmePost{token, "", []PostImg{PostImg{"image", imgLoc}}}
	binData, _ := json.Marshal(params)
	fmt.Println(string(binData))
	req, err := http.NewRequest("Post", postURL,  bytes.NewBuffer(binData))
	req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    defer resp.Body.Close()
    return err
}

func callGiphy(keywords []string) error {
	// group := "20745774"
	file := strings.Join(keywords[:],"+") + ".gif"
	// file = url.QueryEscape(file)
	err := downloadGif(keywords, file)
	if err != nil {
		fmt.Println("ERROR", err)
		return err
	}
	groupmeUrl, err := groupmeImageHost(file)
	defer destroyFile(file)
	if err != nil {
		return err
	}
	err = postGif(groupmeUrl + ".large")
	return err

}



func main() {
	// keywords := []string{"Hitler", "was", "a", "cool", "dude"}
	// callGiphy(keywords)
	

	// err := downloadFile("temp.gif", "http://media3.giphy.com/media/xTiTngCK4tZaoHNlao/giphy.gif")
	// defer destroyFile("temp.gif")
	// fmt.Println("Done")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
