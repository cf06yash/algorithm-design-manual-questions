package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var KEYPAD_BINDINGS map[int][]string

func init() {
	KEYPAD_BINDINGS = make(map[int][]string)
	KEYPAD_BINDINGS[2] = []string{"a", "b", "c"}
	KEYPAD_BINDINGS[3] = []string{"d", "e", "f"}
	KEYPAD_BINDINGS[4] = []string{"g", "h", "i"}
	KEYPAD_BINDINGS[5] = []string{"j", "k", "l"}
	KEYPAD_BINDINGS[6] = []string{"m", "n", "o"}
	KEYPAD_BINDINGS[7] = []string{"p", "q", "r", "s"}
	KEYPAD_BINDINGS[8] = []string{"t", "u", "v"}
	KEYPAD_BINDINGS[9] = []string{"w", "x", "y", "z"}
}

func generate(seq []int, s string, dict map[string]struct{}, idx int, result *[]string) {
	if idx == len(seq) {
		_, ok := dict[s]
		if ok {
			*result = append(*result, s)
		}
		return
	}
	var possibleLetters []string = KEYPAD_BINDINGS[seq[idx]]
	if len(possibleLetters) > 0 {
		generate(seq, s+possibleLetters[0], dict, idx+1, result)
	}
	if len(possibleLetters) > 1 {
		generate(seq, s+possibleLetters[1], dict, idx+1, result)
	}
	if len(possibleLetters) > 2 {
		generate(seq, s+possibleLetters[2], dict, idx+1, result)
	}
	if len(possibleLetters) > 3 {
		generate(seq, s+possibleLetters[3], dict, idx+1, result)
	}
}

func main() {

	/*
		Write a function which, given a sequence of digits 2–9 and a dictionary of n
		words, reports all words described by this sequence when typed in on a standard
		telephone keypad. For the sequence 269 you should return any, box, boy, and
		cow, among other words.
	*/

	dict := map[string]struct{}{
		"boy":     {},
		"box":     {},
		"yellow":  {},
		"cat":     {},
		"mother":  {},
		"batman":  {},
		"serpico": {},
		"charlie": {},
		"shin":    {},
		"ichigo":  {},
		"cow":     {},
	}
	fmt.Println("Please enter sequence:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	seqStr := strings.Split(s, " ")
	var seq []int = make([]int, len(seqStr))
	for i, v := range seqStr {
		seq[i], _ = strconv.Atoi(v)
	}
	fmt.Printf("Entered Sequence is: %v\n", seq)
	result := make([]string, 0)
	generate(seq, "", dict, 0, &result)
	fmt.Printf("Generated Words are: %v\n", result)
}
