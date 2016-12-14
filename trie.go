package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

type Node struct {
	word         string
	wordEndsHere bool
	children     map[byte]*Node
}

func NewNode() *Node {

	children := make(map[byte]*Node)

	return &Node{children:children}
}

func (t *Node) Insert(word string, index int) {

	if index >= len(word) {
		return
	}

	letter := word[index]

	if _, ok := t.children[letter]; !ok {
		t.children[letter] = NewNode()
	}

	if index + 1 == len(word) {

		t.wordEndsHere = true
		t.word = word

	} else {
		t.children[letter].Insert(word, index + 1)
	}
}

func (t *Node) Contains(word string, index int) bool {

	if len(word) == 0 {
		return true
	}

	if len(word) - 1 == index {
		if t.wordEndsHere {
			return true
		} else {
			return false
		}
	}

	letter := word[index]

	if _, ok := t.children[letter]; !ok {
		return false
	} else {
		return t.children[letter].Contains(word, index + 1)
	}
}

func (t *Node) GetAllWithPrefix(prefix string, index int) []string {

	results := make([]string, 0)

	if index >= len(prefix) {

		for _, child := range t.children {

			if child.wordEndsHere {

				results = append(results, child.word)

			}

			for _, item := range child.GetAllWithPrefix(prefix, index) {
				results = append(results, item)
			}
		}

	} else if _, ok := t.children[prefix[index]]; ok {

		for _, item := range t.children[prefix[index]].GetAllWithPrefix(prefix, index + 1) {

			results = append(results, item)
		}
	}

	return results
}

func main() {

	contents, err := ioutil.ReadFile("/usr/share/dict/words")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	words := strings.Split(string(contents), "\n")

	trie := NewNode()

	for _, word := range words {

		trie.Insert(word, 0)
	}

	for _, word := range trie.GetAllWithPrefix("z", 0) {
		fmt.Println(word)
	}
}