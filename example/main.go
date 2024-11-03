package main

import (
	ua "github.com/liuyuhe666/random-ua"
	"log"
)

func main() {
	// RandomUserAgent generates a random DESKTOP browser user-agent
	log.Printf("RandomUserAgent: %s\n", ua.RandomUserAgent())
	// RandomMobileUserAgent generates a random MOBILE browser user-agent
	log.Printf("RandomMobileUserAgent: %s\n", ua.RandomMobileUserAgent())
}
