package main

import (
	"fmt"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var trie = NewNode()

func handler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("q")
	matches := trie.KeysWithPrefix(query)

	results, err := json.Marshal(matches)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, string(results))
}

func populateTrie(words []string) {

	for _, word := range words {
		trie.Insert(word)
	}
}

func main() {

	query_file := os.Getenv("QUERY_FILE")

	if query_file == "" {
		fmt.Println("Set $QUERY_FILE to path to list of search queries")
		os.Exit(1)
	}

	contents, err := ioutil.ReadFile("/usr/share/dict/words")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	populateTrie(strings.Split(string(contents), "\n"))

	http.HandleFunc("/autocomplete", handler)
	http.ListenAndServe(":8080", nil)
}
