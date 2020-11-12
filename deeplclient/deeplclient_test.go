package deeplclient

import (
	"log"
	"os"
	"testing"
)

func TestDeeplClient_TranslateText(t *testing.T) {
	apikey := os.Getenv("DEEPL_APIKEY")
	client := DeeplClient{Apikey: apikey}
	test := "テスト"
	res, err := client.TranslateText(test, "", "EN")
	if err != nil {
		log.Fatalf("Error happened %s", err)
	}
	text := res.Translations[0].Text
	detectedlan := res.Translations[0].DetectedSourceLanguage

	if text != "test" {
		log.Fatalf("The result is not what you expected. the result was %s", text)
	}
	if detectedlan != "JA" {
		log.Fatalf("The detected language wat not expected. got=%s", detectedlan)
	}
}
