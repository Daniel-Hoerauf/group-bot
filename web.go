package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"net/http"
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
	message := string(body[:])
	response := GroupmeContent{}
	json.Unmarshal(bin, &response)
	fmt.Println(response)
}

func postGif(imgLoc string) (error) {
	postURL := "https://api.groupme.com/v3/bots/post"
	token := os.Getenv("GROUPME_BOT_TOKEN")
	params := GroupmePost{token, "Hitler was an ok dude and a gifted artist", []PostImg{PostImg{"image", imgLoc}}}
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
