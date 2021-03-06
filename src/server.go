package main

import (
	"fmt"
	"os"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var trie = NewTrie()

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
	query_url := os.Getenv("QUERY_URL")
	port := os.Getenv("PORT")

	if port == "" || rune(port[0]) != ':' {
		port = ":80"
	}

	var contents []byte
	var resp *http.Response
	var err error

	if query_file != "" {

		contents, err = ioutil.ReadFile(query_file)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}


	} else if query_url != "" {

		resp, err = http.Get(query_url)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer resp.Body.Close()

		contents, err = ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	} else {
		fmt.Println("Set $QUERY_FILE or $QUERY_URL to path or URL to list of search queries")
		os.Exit(1)
	}


	fmt.Println("Populating the trie")
	populateTrie(strings.Split(string(contents), "\n"))

	fmt.Println("Listening...")
	http.HandleFunc("/autocomplete", handler)
	http.ListenAndServe(port, nil)
}
