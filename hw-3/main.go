package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func textToSlice(s string) []string {
	//find words in text
	inputStrings := regexp.MustCompile("\\w+").FindAllString(s, -1)

	//format strings to lower case
	for i := 0; i < len(inputStrings); i++ {
		inputStrings[i] = strings.ToLower(inputStrings[i])
	}

	return inputStrings
}

func mostPopularWordsInText(inputStrings []string) {
	inputStringsInMap := make(map[string]int)

	//add strings to map
	for i := 0; i < len(inputStrings); i++ {
		inputStringsInMap[inputStrings[i]]++
	}

	values := make([]int, 0, len(inputStringsInMap))
	for _, val := range inputStringsInMap {
		values = append(values, val)
	}

	//sort by values of map
	sort.Ints(values)
	pointBreak := 0

	// print top10 popular words in text
	for i := len(values) - 1; i > 0 && pointBreak < 10; i-- {
		for key, val := range inputStringsInMap {
			if values[i] == val && pointBreak < 10 {
				pointBreak++
				fmt.Printf("%s %d\n", key, val)
			}
		}
	}
}

func main() {

	s := `When a humble bard
	Graced a ride along
	With Geralt of Rivia
	Along came this song
	From when the White Wolf fought
	A silver-tongued devil
	His army of elves
	At his hooves did they revel
	They came after me
	With masterful deceit
	Broke down my lute
	And they kicked in my teeth
	While the devil's horns
	Minced our tender meat
	And so cried the Witcher
	"He can't be bleat"
	Toss a coin to your Witcher
	O' Valley of Plenty
	O' Valley of Plenty, oh
	Toss a coin to Your Witcher
	O' Valley of Plenty
	At the edge of the world
	Fight the mighty horde
	That bashes and breaks you
	And brings you to mourn
	He thrust every elf
	Far back on the shelf
	High up on the mountain
	From whence it came
	He wiped out your pest
	Got kicked in his chest
	He's a friend of humanity
	So give him the rest
	That's my epic tale:
	Our champion prevailed
	Defeated the villain
	Now pour him some ale
	Toss a coin to your Witcher
	O' Valley of Plenty
	O' Valley of Plenty, oh
	Toss a coin to your Witcher
	A friend of humanity
	Toss a coin to your Witcher
	O' Valley of Plenty
	O' Valley of Plenty, oh
	Toss a coin to your Witcher
	A friend of humanity
	Toss a coin to your Witcher
	O' Valley of Plenty
	O' Valley of Plenty, a-oh
	Toss a coin to your Witcher
	A friend of humanity`
	mostPopularWordsInText(textToSlice(s))

}
