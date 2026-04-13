package main

import (
	"flag"
	"fmt"
)

type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"de": "Hallo Welt",
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"he": "שלום עולם",
	"ur": "ہیلو دنیا",
	"vi": "Xin chào Thế Giới",
}

func main() {
	var lang string

	flag.StringVar(&lang,
		"lang",
		"en",
		"The required language, e.g. en, fr, ur...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

func greet(l language) string {
	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
