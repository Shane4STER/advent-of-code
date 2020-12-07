package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type contents struct {
	count int
	bag   *bag
}

type bag struct {
	colour   string
	contents []contents
}

func main() {
	var allBags = make(map[string]*bag)

	scanner := bufio.NewScanner(openStdinOrFile())

	for scanner.Scan() {
		var currentBag *bag
		line := scanner.Text()

		containerColour, containerContents := parseLine(line)

		if cb, exists := allBags[containerColour]; !exists {
			currentBag = &bag{containerColour, make([]contents, 0)}
			allBags[containerColour] = currentBag
		} else {
			currentBag = cb
		}

		for _, s := range containerContents {
			var targetBag *bag
			s = strings.TrimSpace(s)
			count, err := strconv.Atoi(strings.Split(s, " ")[0])
			if err != nil {
				continue
			}
			colour := strings.Join(strings.Split(s, " ")[1:3], " ")
			if tb, exists := allBags[colour]; !exists {
				targetBag = &bag{colour, make([]contents, 0)}
				allBags[colour] = targetBag
			} else {
				targetBag = tb
			}
			currentBag.contents = append(currentBag.contents, contents{count, targetBag})
		}
	}
	fmt.Printf("%v\n", findBagNumber(allBags, "shiny gold"))
	fmt.Printf("%v\n", allBags["shiny gold"].countSubBags())
}

func findBagNumber(allBags map[string]*bag, targetBagColour string) int {
	var count int
	targetBag, exists := allBags[targetBagColour]
	if !exists {
		return 0
	}

	for _, bag := range allBags {
		if bag.contains(targetBag, true) {
			count++
		}
	}
	return count
}

func (b bag) countSubBags() int {
	var subbags int
	for _, content := range b.contents {
		childSubBags := content.bag.countSubBags()
		subbags += content.count * (childSubBags + 1)
	}
	return subbags
}

func (b bag) contains(target *bag, recurse bool) bool {
	for _, content := range b.contents {
		if content.bag == target {
			return true
		}
	}
	if recurse {
		for _, content := range b.contents {
			if content.bag.contains(target, true) {
				return true
			}
		}
	}
	return false
}

func parseLine(line string) (string, []string) {
	containerColour := strings.Join(strings.Split(line, " ")[:2], " ")

	containerContents := strings.Split(strings.Join(strings.Split(line, " ")[4:], " "), ",")

	return containerColour, containerContents
}

func openStdinOrFile() io.Reader {
	var err error
	r := os.Stdin
	if len(os.Args) > 1 {
		r, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	return r
}
