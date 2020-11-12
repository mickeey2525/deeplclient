package deeplclient

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type deeplResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

// DeeplClient Constructor
type DeeplClient struct {
	Apikey string
}

func NewClient(key string) *DeeplClient {
	dlc := DeeplClient{}
	dlc.Apikey = key
	return &dlc
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("can't close body!!", err)
		}
	}()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func (dc *DeeplClient) TranslateText(texttarget, srclang, targetlang string) (*deeplResponse, error) {

	textList := splitText(texttarget)
	data := url.Values{
		"auth_key":    {dc.Apikey},
		"source_lang": {srclang},
		"target_lang": {targetlang},
	}

	for _, v := range textList {
		data.Add("text", v)
	}
	resp, err := http.PostForm("https://api.deepl.com/v2/translate", data)
	if err != nil {
		log.Fatalf("an error happed during calling the function %ss", err)
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("can't close body!!", err)
		}
	}()
	if resp.StatusCode != 200 {
		log.Fatalf("Something Went wrong. The statsu code is %+v, and responce body is %+v", resp.StatusCode, resp.Body)
	}
	var res deeplResponse
	if err := decodeBody(resp, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func splitText(text string) []string {
	textList := strings.Split(strings.Replace(text, "\r\n", "\n", -1), "\n")
	return textList
}

func (dc *DeeplClient) TranslateFile(filepath, srclang, targetlang string) (*deeplResponse, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("An Error happed during opening a file: %s", err)
	}

	textList := splitText(string(file))
	data := url.Values{
		"auth_key":    {dc.Apikey},
		"source_lang": {srclang},
		"target_lang": {targetlang},
	}

	for _, v := range textList {
		data.Add("text", v)
	}

	resp, err := http.PostForm("https://api.deepl.com/v2/translate", data)
	if err != nil {
		log.Fatalf("an error happed during calling the function %s", err)
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("can't close http connection!!", err)
		}
	}()
	if resp.StatusCode != 200 {
		log.Fatalf("Something Went wrong. The statsu code is %+v, and responce body is %+v", resp.StatusCode, resp.Body)
	}
	var res deeplResponse
	if err := decodeBody(resp, &res); err != nil {
		log.Fatalf("Something went wrong during decode the reponce body. The error message is %s", err)
		return nil, err
	}
	return &res, nil
}
