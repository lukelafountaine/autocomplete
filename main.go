package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

func main() {

	contents, err := ioutil.ReadFile("/usr/share/dict/words")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	words := strings.Split(string(contents), "\n")

	trie := NewNode()

	// test Insert
	for _, word := range words {
		trie.Insert(word)
	}

	// test Contains
	for _, word := range words {
		if !trie.Contains(word) {
			fmt.Println(word)
		}
	}

	// test AllKeys
	if len(trie.AllKeys()) != len(words) {
		fmt.Println(len(trie.AllKeys()), len(words))
	}

	// test KeysWithPrefix
	fmt.Println(trie.KeysWithPrefix("artifact"))
}
