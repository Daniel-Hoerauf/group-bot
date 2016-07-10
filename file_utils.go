package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func buildUrl(keywords []string) (giphyURL string) {
	args := strings.Join(keywords[:], "+")
	url := fmt.Sprintf("http://api.giphy.com/v1/gifs/translate?s=%s&api_key=dc6zaTOxFJmzC&rating=r", args)
	return url
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func downloadGif(keywords []string) (*bytes.Buffer, error) {
	url := buildUrl(keywords)
	fmt.Println(url)

	resp := GiphyResponse{}
	err := getJson(url, &resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	gif, err := http.Get(resp.Data.Images.Down.Url)
	if err != nil {
		return nil, err
	}
	defer gif.Body.Close()
	buf, err := ioutil.ReadAll(gif.Body)

	return bytes.NewBuffer(buf), err

}

func groupmeImageHost(gif *bytes.Buffer) (img string, err error) {

	// token := os.Getenv("GROUPME_ACCESS_TOKEN")
	token := secrets.Token
	url := fmt.Sprintf("https://image.groupme.com/pictures?access_token=%s", token)

	req, err := http.NewRequest("POST", url, gif)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	bin, _ := ioutil.ReadAll(resp.Body)

	response := GroupmeResponse{}
	json.Unmarshal(bin, &response)
	return response.Payld.Url, nil

}
