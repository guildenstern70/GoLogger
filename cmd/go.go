/*
 * Project GoLogger
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

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
		fmt.Printf("%s > %s%s%s\n", CurDateTime(), alerter(), sentence(), detailer())
		time.Sleep(sleep * time.Millisecond)
	}
}

func sentence() string {
	return fmt.Sprintf("%s %s %s %s",
		GetNoun(true),
		GetNoun(false),
		GetVerb(true),
		GetNoun(false))
}

func detail() string {
	return fmt.Sprintf("%s %s %s",
		GetNoun(false),
		GetVerb(true),
		GetNoun(false))
}

func alerter() string {
	var dice = rand.Intn(12)
	if dice == 7 {
		return "ERROR: "
	} else if dice == 8 || dice == 9 {
		return "WARNING: "
	}
	return ""
}

func detailer() string {
	var dice = rand.Intn(12)
	var comma = dice%2 == 0
	if dice%3 == 0 {
		return GetConjunctions(comma) + detail()
	}
	return ""
}
