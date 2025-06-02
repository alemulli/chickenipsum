package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

/*
 * Copyright (c) 2009, Mathew Tinsley (tinsley@tinsology.net)
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *    * Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *    * Redistributions in binary form must reproduce the above copyright
 *      notice, this list of conditions and the following disclaimer in the
 *      documentation and/or other materials provided with the distribution.
 *    * Neither the name of the organization nor the
 *      names of its contributors may be used to endorse or promote products
 *      derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY MATHEW TINSLEY ''AS IS'' AND ANY
 * EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL <copyright holder> BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

// LoremIpsum represents the chickenipsum generator
type LoremIpsum struct{}

// Constants for sentence generation
const (
	WordsPerSentenceAvg = 8
	WordsPerSentenceStd = 4
)

// List of possible chicken-themed words
var words = []string{
	// Normal chicken words
	"disco",
	"chicken",
	"hen",
	"chick",
	"coop critter",
	"feathered friend",
	"sky floof",
	"flapper",
	"sky bean",
	"cluck",
	"chickengorino",
	"extremely cuuuuuute",
	"bawk",
	"peck",
	"peckin' peck",
	"scratch-scratch",
	"featherfluff",
	"henhouse",
	"eggcellent",
	"chickadee",
	"coop commander",
	"flapjack",
	"nest snuggler",
	"egg-layer",
	"broody",
	"cock-a-doodle-doo",
	"roost ruler",
	"wing-ding",
	"chirpy chirp",
	"dust bather",
	"rooster booster",
	"cluckster",
	"fluffy nugget",
	"peep peep",
	"waddle waddle",
	"feathery fiend",
	"clucklehead",
	"egg dropper",
	"chick magnet",
	"plucky",
	"poultrygeist",
	"yolk yolker",
	"nest napper",
	"scramble master",
	"beak bonker",
	"flap flap",
	"coop-tastic",
	"feather duster",
	"squawkbox",
	"hen-tertainer",
	"eggstreme",
	"scratch patrol",
	"cackle queen",
	"fowl friend",
	"biddy bop",
	"egg machine",
	"chick-chick hooray",
	"fluff nugget",
	"bawkbag",
	"pecklehead",
	"clucklebutt",
	"egg wobbler",
	"feather floof",
	"scritch scratcher",
	"chirp muffin",
	"flapflap",
	"beaky mcpeckface",
	"nugget noodle",
	"waddle puff",
	"cluckmuffin",
	"egg sneezer",
	"squawk lump",
	"biddy bonkers",
	"yolk yelper",
	"flufflebawk",
	"scramble puff",
	"perch potato",
	"coop goblin",
	"nest beast",
	"peep squeak",
	"wingy dingy",
	"drumstick doofus",
	"plop pecker",
	"feather dingle",
	"broodlesnoot",
	"cheep chunk",
	"bawklet",
	"egg plopper",
	"snood snoozer",
	"chickeny chonk",
	"henwiggle",
	"beak blinker",
	"cluckerdoodle",
	"tiny squawker",
	"pecky mcflap",
	"fluffstomp",
	"eggbubble",
	"peep nugget",
	"tiny floof",
	"chirple bean",
	"mini bawk",
	"fluff pebble",
	"snoozle chick",
	"pufflet",
	"beak bean",
	"chirp chunk",
	"snuggle peeper",
	"wobble floof",
	"egglet",
	"cheep muffin",
	"wee squawker",
	"peepadoop",
	"chirpy blob",
	"floof squeak",
	"snack-sized squawk",
	"fuzzy noodle",
	"puddle peep",
	"tiny plopper",
	"chicklet chonk",
	"pocket clucker",
	"boop beaker",
	"scritchlet",
	"fuzz niblet",
	"bitty biddy",
	"chirp chub",
	"cheep cheepums",
	"wiggle fluff",
	"mini pluff",
	"peepadoo",
	"floofnug",
	"baby bawkbean",
	"snuggle nugget",
	"beakling",
	"chirp squish",
	"feather bean",
	"yolkling",
	"squeakle puff",
	"bokzilla roar",
	"squawkle",
	"flap squawk",
	"cheep cheep",
	"buh-gawk",
	"rawwk",
	"bawkawk",
	"cluckle",
	"bok bok",
	"bok-bok-bok",
	"big boi",
	"tiny",
	"wobble-tastic",
	"bonkers mode",
	"snuggle-powered",
	"derp-certified",
	"zoom-enabled",
	"boop-happy",
	"clumsy deluxe",
	"crash enthusiast",
	"hovering awkwardly",
	"flailing gently",
	"scoot-enabled",
	"gobsmacked and flappy",
	"squibble certified",
	"mildly feral",
	"flop mode engaged",
	"shenanigan-prone",
	"sky nugget",
	// Disco chicken words
	"boogie bawk",
	"funky feather",
	"cluckin’ groovy",
	"peck-a-beat",
	"disco chick",
	"feather hustle",
	"dj eggscrambler",
	"saturday night cluck",
	"groove rooster",
	"fowl on the floor",
	"bawkstep",
	"feather flares",
	"cluck 'n' roll",
	"the yolk shimmy",
	"hustle hen",
	"dancin’ drumstick",
	"bawk and bounce",
	"egg shake",
	"chicka-chicka-boom",
	"boogie coop",
	"scramble slide",
	"flap funk",
	"poultry party",
	"hen-hop",
	"glowstick clucker",
	"cluckadelic",
	"yolk-a-doodle-doo",
	"beegees bantam",
	"feather boogie",
	"funky coop beat",
	"mirror ball biddy",
	"disco nest",
	"groove eggspress",
	"the cluckstep",
	"bass beaker",
	"boogiedoodle",
	"peep pop",
	"shake-a-tailfeather",
	"roller rooster",
	"jive turkey’s cousin",
	// Chicken names
	"Speckles",
	"Chloe",
	"Big Chonkers",
	"Laquisha",
	"Shanaynay",
	"Pearl",
	"La Blanca",
	"Buffy",
	"Butterscotch",
	"Red",
	"Ruby",
	"Hawk-Eye",
	"Precious",
}

// Generate creates "Lorem ipsum" style words with chicken-themed content
func (l *LoremIpsum) Generate(numWords int) string {
	if numWords <= 0 {
		numWords = 100 // Default value
	}

	// Start with first two words
	generatedWords := []string{words[0], words[1]}
	remaining := numWords - 2

	// Add remaining words
	for i := 0; i < remaining; i++ {
		position := rand.Intn(len(words))
		word := words[position]

		// Avoid repeating the same word consecutively
		if len(generatedWords) > 0 && generatedWords[len(generatedWords)-1] == word {
			i--
			continue
		}

		generatedWords = append(generatedWords, word)
	}

	sentences := []string{}
	current := 0
	wordsLeft := len(generatedWords)

	// Create sentences from the generated words
	for wordsLeft > 0 {
		sentenceLength := l.getRandomSentenceLength()

		if wordsLeft < sentenceLength {
			sentenceLength = wordsLeft
		}

		if sentenceLength < 1 {
			break
		}

		sentence := generatedWords[current : current+sentenceLength]
		sentence = l.punctuate(sentence)
		current += sentenceLength
		wordsLeft -= sentenceLength
		sentences = append(sentences, strings.Join(sentence, " "))
	}

	return strings.Join(sentences, " ")
}

// punctuate inserts commas and periods in the given sentence
func (l *LoremIpsum) punctuate(sentence []string) []string {
	wordLength := len(sentence)

	// Handle empty sentences or single word sentences
	if wordLength == 0 {
		return sentence
	}

	// End the sentence with a period
	sentence[wordLength-1] = sentence[wordLength-1] + "."

	if wordLength < 4 {
		// Capitalize the first letter of the first word
		if len(sentence[0]) > 0 {
			sentence[0] = strings.ToUpper(sentence[0][:1]) + sentence[0][1:]
		}
		return sentence
	}

	numCommas := l.getRandomCommaCount(wordLength)

	for i := 0; i <= numCommas; i++ {
		position := int(math.Round(float64(i) * float64(wordLength) / float64(numCommas+1)))

		if position < (wordLength-1) && position > 0 {
			// Add the comma
			sentence[position] = sentence[position] + ","
		}
	}

	// Capitalize the first letter of the first word
	if len(sentence[0]) > 0 {
		sentence[0] = strings.ToUpper(sentence[0][:1]) + sentence[0][1:]
	}

	return sentence
}

// getRandomCommaCount produces a random number of commas
func (l *LoremIpsum) getRandomCommaCount(wordLength int) int {
	base := 6.0 // Arbitrary

	average := math.Log(float64(wordLength)) / math.Log(base)
	standardDeviation := average / base

	return int(math.Round(l.gaussMS(average, standardDeviation)))
}

// getRandomSentenceLength produces a random sentence length
func (l *LoremIpsum) getRandomSentenceLength() int {
	length := int(math.Round(l.gaussMS(WordsPerSentenceAvg, WordsPerSentenceStd)))
	if length < 1 {
		return 1 // Ensure we always have at least one word per sentence
	}
	return length
}

// gauss produces a random number
func (l *LoremIpsum) gauss() float64 {
	return (rand.Float64()*2 - 1) +
		(rand.Float64()*2 - 1) +
		(rand.Float64()*2 - 1)
}

// gaussMS produces a random number with Gaussian distribution
func (l *LoremIpsum) gaussMS(mean, standardDeviation float64) float64 {
	return l.gauss()*standardDeviation + mean
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Parse command line arguments
	numWords := flag.Int("words", 100, "Number of words to generate")
	flag.Parse()

	// Create lorem ipsum generator
	lorem := &LoremIpsum{}

	// Generate and print the chicken-themed lorem ipsum text
	fmt.Println(lorem.Generate(*numWords))
}
