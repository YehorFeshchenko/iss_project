package indexer

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Node struct {
	Word string
	Freq []int
	Next *Node
}

func AddNodeToList(head *Node, word string, fileIndex int) *Node {
	newNode := &Node{Word: word}

	if head == nil || head.Word > word {
		newNode.Freq = make([]int, fileIndex+1)
		newNode.Freq[fileIndex] = 1
		newNode.Next = head
		return newNode
	}

	current := head
	var prev *Node
	for current != nil && current.Word < word {
		prev = current
		current = current.Next
	}

	if current != nil && current.Word == word {
		if len(current.Freq) <= fileIndex {
			newFreq := make([]int, fileIndex+1)
			copy(newFreq, current.Freq)
			current.Freq = newFreq
		}
		current.Freq[fileIndex] += 1
	} else {
		newNode.Freq = make([]int, fileIndex+1)
		newNode.Freq[fileIndex] = 1
		newNode.Next = current
		if prev != nil {
			prev.Next = newNode
		} else {
			head = newNode
		}
	}

	return head
}

func RemovePunctuationFromStartEnd(word string) string {
	word = strings.ToLower(word)

	start := 0
	for ; start < len(word) && !unicode.IsLetter(rune(word[start])) && !unicode.IsNumber(rune(word[start])); start++ {
	}

	end := len(word) - 1
	for ; end >= 0 && !unicode.IsLetter(rune(word[end])) && !unicode.IsNumber(rune(word[end])); end-- {
	}

	if start > end {
		return ""
	}

	return word[start : end+1]
}

func IndexFiles(files []string, indexFilename, resultFilename string) error {
	var indexHead *Node
	var filenames []string

	for fileIndex, fileName := range files {
		filenames = append(filenames, fileName)
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			for _, word := range words {
				cleanWord := RemovePunctuationFromStartEnd(word)
				if cleanWord == "" {
					continue
				}
				indexHead = AddNodeToList(indexHead, cleanWord, fileIndex)
			}
		}

		file.Close()
	}

	return writeResults(indexHead, filenames, indexFilename, resultFilename)
}

func writeResults(indexHead *Node, filenames []string, indexFilename, filenamesFilename string) error {
	file, err := os.Create("/app/data/" + filenamesFilename + ".txt")
	if err != nil {
		return err
	}
	for _, name := range filenames {
		if _, err := file.WriteString(name + "\n"); err != nil {
			file.Close()
			return err
		}
	}
	file.Close()

	indexFile, err := os.Create("/app/data/" + indexFilename + ".txt")
	if err != nil {
		return err
	}
	defer indexFile.Close()

	for current := indexHead; current != nil; current = current.Next {
		for len(current.Freq) < len(filenames) {
			current.Freq = append(current.Freq, 0)
		}

		freqStr := make([]string, len(current.Freq))
		for i, freq := range current.Freq {
			freqStr[i] = strconv.Itoa(freq)
		}
		if _, err := indexFile.WriteString(current.Word + " " + strings.Join(freqStr, " ") + "\n"); err != nil {
			return err
		}
	}

	return nil
}
