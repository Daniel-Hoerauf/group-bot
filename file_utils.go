package main

import (
  "os"
  "encoding/json"
  "io"
  "io/ioutil"
  "bytes"
  "fmt"
  "net/http"
  "net/url"
  "strings"
)

func downloadFile(filepath string, url string) (err error) {

  // Create the file
  out, err := os.Create(filepath)
  if err != nil  {
    return err
  }
  defer out.Close()

  // Get the data
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // Writer the body to file
  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    return err
  }

  return nil
}

func buildUrl(keywords []string) (giphyURL string) {
	args := strings.Join(keywords[:],"+")
	query := url.QueryEscape(args)
	url := fmt.Sprintf("http://api.giphy.com/v1/gifs/translate?s=%s&api_key=dc6zaTOxFJmzC", query)
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

func downloadGif(keywords []string, file string) error {
	url := buildUrl(keywords)
	fmt.Println(url)

	resp := GiphyResponse{}
	err := getJson(url, &resp)
	// resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = downloadFile(file, resp.Data.Images.Down.Url)
	// err := downloadFile(file, url)
	return err

}

func destroyFile(filepath string) error {
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File", filepath, "destroyed")
	return err
}

func groupmeImageHost(file string) (img string, err error) {
  buf, err := ioutil.ReadFile(file)
  if err != nil {
    return "", err
  }
  // fi := string(buf)
  // v := url.Values{}
  // v.Add("file", fi)

  token := os.Getenv("GROUPME_ACCESS_TOKEN")
  url := fmt.Sprintf("https://image.groupme.com/pictures?access_token=%s", token)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(buf))

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