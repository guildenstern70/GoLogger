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
	"Wednesday",
	"Thursday",
	"yesterday",
	"today",
	"sentiment",
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
	"new",
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

var verbs = []string{
	"associate",
	"fear",
	"spoil",
	"exercise",
	"guarantee",
	"tap",
	"release",
	"appeal",
	"discharge",
	"swing",
	"circulate",
	"submit",
	"house",
	"confer",
	"cast",
	"shut",
	"rip",
	"educate",
	"dare",
	"disagree",
	"drop",
	"strip",
	"age",
	"wrap",
	"choose",
	"classify",
	"fall",
	"shout",
	"describe",
	"persist",
}

var conjunctions = []string{
	"but",
	"nonetheless",
	"if",
	"on the contrary",
	"at the end",
	"finally",
	"that",
	"so that",
	"but",
	"but",
	"if",
}

func GetNoun(capitalize bool) string {
	var idx = rand.Intn(len(nouns))
	var noun = nouns[idx]
	if capitalize {
		return strings.Title(noun)
	}
	return noun
}

func GetVerb(thirdperson bool) string {
	var idx = rand.Intn(len(verbs))
	var verb = verbs[idx]
	if thirdperson {
		verb += "s"
	}
	return verb
}

func GetConjunctions(commabefore bool) string {
	var idx = rand.Intn(len(conjunctions))
	var conjunction = conjunctions[idx] + " "
	if commabefore {
		return ", " + conjunction
	}
	return " " + conjunction
}
