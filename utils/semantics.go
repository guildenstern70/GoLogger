/*
 * Project GoLogger
 * Copyright (c) Alessio Saltarin 2021.
 * Licensed under MIT Licence - See LICENSE
 */

package utils

import (
	"math/rand"
	"strings"
)

var nouns = []string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
	"conversation",
	"road",
	"reflection",
	"anxiety",
	"funeral",
	"secretary",
	"attitude",
	"setting",
	"bath",
	"variety",
	"user",
	"construction",
	"art",
	"thanks",
	"presentation",
	"world",
	"possibility",
	"alcohol",
	"student",
	"error",
	"sector",
	"hat",
	"painting",
	"agency",
	"investment",
	"warning",
	"player",
	"manager",
	"mom",
	"coffee",
	"appointment",
	"road",
	"initiative",
	"love",
	"thing",
	"opportunity",
	"loss",
	"user",
	"efficiency",
	"theory",
	"contract",
	"city",
	"wedding",
	"indication",
}

func GetNoun(capitalize bool) string {
	var idx = rand.Intn(len(nouns))
	var noun = nouns[idx]
	if capitalize {
		return strings.Title(noun)
	}
	return noun
}
