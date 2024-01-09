package query

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	Word string
	Freq []int
	Next *Node
}

func getIndexHead(filename string) *Node {
	head := &Node{Word: ""}
	current := head

	file, err := os.Open(filename + ".txt")
	if err != nil {
		fmt.Println("Error opening index file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		word := parts[0]
		var frequencies []int
		for _, freqStr := range parts[1:] {
			freq, err := strconv.Atoi(freqStr)
			if err != nil {
				fmt.Println("Error parsing frequency:", err)
				continue
			}
			frequencies = append(frequencies, freq)
		}

		newNode := &Node{Word: word, Freq: frequencies}
		current.Next = newNode
		current = newNode
	}

	return head.Next
}

func getFilenames(filename string) []string {
	var filenames []string

	file, err := os.Open(filename + ".txt")
	if err != nil {
		fmt.Println("Error opening filenames file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		filenames = append(filenames, scanner.Text())
	}

	return filenames
}

func ExecuteQuery(word, indexFilename, filenamesFilename string) {
	indexHead := getIndexHead(indexFilename)
	filenames := getFilenames(filenamesFilename)

	var freqFilenamePairs []struct {
		Freq     int
		Filename string
	}

	for current := indexHead; current != nil; current = current.Next {
		if current.Word == word {
			for i, freq := range current.Freq {
				if freq > 0 {
					freqFilenamePairs = append(freqFilenamePairs, struct {
						Freq     int
						Filename string
					}{freq, filenames[i]})
				}
			}
			break
		}
	}

	if len(freqFilenamePairs) == 0 {
		fmt.Println("No matches found in index")
		return
	}

	sort.Slice(freqFilenamePairs, func(i, j int) bool {
		return freqFilenamePairs[i].Freq > freqFilenamePairs[j].Freq
	})

	for _, pair := range freqFilenamePairs {
		fmt.Printf("%d %s\n", pair.Freq, pair.Filename)
	}
}
