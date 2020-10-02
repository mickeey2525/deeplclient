package deeplclient

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type deeplResponce struct {
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

func (dc *DeeplClient) TranslateText(texttarget, srclang, targetlang string) (*deeplResponce, error){
	data := url.Values{
		"auth_key":    {dc.Apikey},
		"text":        {texttarget},
		"source_lang": {srclang},
		"target_lang": {targetlang},
	}
	resp, err := http.PostForm("https://api.deepl.com/v2/translate", data)
	if err != nil {
		log.Fatalf("%s",err)
		return nil, err
	}
	defer resp.Body.Close()
	var res deeplResponce
	json.NewDecoder(resp.Body).Decode(&res)
	return &res,nil
}
