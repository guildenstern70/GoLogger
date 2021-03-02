package main

import (
	"fmt"
	. "github.com/guildenstern70/gologger/utils"
	"math/rand"
	"time"
)

// VERSION OF GO-LOGGER
const VERSION = "0.1"

func main() {
	fmt.Printf("GoLogger v.%s\n\n", VERSION)
	time.Sleep(1 * time.Second)
	var k = 1
	for k < 100 {
		k++
		var sleep = time.Duration(rand.Intn(1500) + 400)
		fmt.Printf("%s > %s %s \n", CurDateTime(), GetNoun(true), GetNoun(false))
		time.Sleep(sleep * time.Millisecond)
	}
}
