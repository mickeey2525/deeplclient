package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mickeey2525/deeplclient/deeplclient"
)

func main() {
	apikey := os.Getenv("DEEPL_APIKEY")
	if apikey == "" {
		log.Fatalf("DEEPL_APIKEY is not set. Please set you apikey as enviroment Variable")
	}
	var (
		text       = flag.String("text", "", "You need to specify the target text")
		filepath   = flag.String("filepath", "", "The text filepath you want to translate")
		srcLang    = flag.String("srclang", "JA", "src language")
		targetLang = flag.String("targetlang", "EN", "target language")
	)

	flag.Parse()
	if *text != "" && *filepath == "" {
		dcl := deeplclient.NewClient(apikey)
		translated, err := dcl.TranslateText(*text, *srcLang, *targetLang)
		if err != nil {
			log.Fatalf("%s", err)
		}
		src, err := json.Marshal(translated)
		if err != nil {
			log.Fatalf("Error happed during convert to JSON")
		}
		fmt.Printf("%s", src)
	}
	if *filepath != "" && *text == "" {
		dcl := deeplclient.NewClient(apikey)
		translated, err := dcl.TranslateFile(*filepath, *srcLang, *targetLang)
		if err != nil {
			log.Fatalf("%s", err)
		}
		src, err := json.Marshal(translated)
		if err != nil {
			log.Fatalf("Error happed during convert to JSON")
		}
		fmt.Printf("%s", src)
	}
	if *filepath == "" && *text == "" {
		fmt.Println("You must specify -text or -filepath. For details, please run deeplclient -h")
	}
}
