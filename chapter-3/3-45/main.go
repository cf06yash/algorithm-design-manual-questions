package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Clean(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	return re.ReplaceAllString(s, "")
}

func main() {
	/*
		Give an algorithm for finding an ordered word pair (e.g. “New York”) occurring with the greatest frequency in a given webpage. Which data structures would you use? Optimize both time and space.
	*/
	var webpage string = "The quick brown fox jumps over the lazy dog. The quick brown fox is very quick and the dog is very lazy. Fox and dog are often found together in stories. The quick brown fox is a common phrase used for typing practice. The lazy dog is not as famous as the quick brown fox."
	cleanedWebPage := Clean(webpage)

	var stringArray []string = strings.Fields(cleanedWebPage)
	var mp map[string]int = make(map[string]int)
	for i := 0; i < len(stringArray)-1; i++ {
		pair := strings.Trim(stringArray[i], " ") + " " + strings.Trim(stringArray[i+1], " ")
		c, ok := mp[pair]
		if ok {
			mp[pair] = c + 1
		} else {
			mp[pair] = 1
		}
	}

	var maxPair string = ""
	var maxCount int = -1

	for pair, value := range mp {
		if value > maxCount {
			maxPair = pair
			maxCount = value
		}
	}

	fmt.Printf("MaxPair: %s and its frequency: %d\n", maxPair, maxCount)
}
