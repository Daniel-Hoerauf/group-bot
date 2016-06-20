package main

import (
  "os"
  "encoding/json"
  "io"
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
	err = downloadFile("temp.gif", resp.Data.Images.Original.Url)
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
