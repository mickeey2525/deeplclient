package main

import (
	"fmt"
	"github.com/mickeey2525/deepl-client/deeplclient"
	"log"
	"os"
)

func main() {
	apikey := os.Args[0]
	dcl := deeplclient.NewClient(apikey)
	a, err:= dcl.TranslateText("テスト", "JA", "EN")
	if err != nil {
		log.Fatalf("%s",err)
	}
	fmt.Println(a)
}
