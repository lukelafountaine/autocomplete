package main

import (
	"sort"
)

type Trie struct {
	word     string
	children map[byte]*Trie
	count    int
}

// type to sort by word count
type ByDescendingCount []*Trie

func (t ByDescendingCount) Len() int {
	return len(t)
}
func (t ByDescendingCount) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t ByDescendingCount) Less(i, j int) bool {
	return t[i].count > t[j].count
}

func NewTrie() *Trie {
	children := make(map[byte]*Trie)
	return &Trie{children:children}
}

func (t *Trie) Insert(word string) {
	t.insertHelper(word, 0)
}

func (t *Trie) insertHelper(word string, index int) *Trie {

	if t == nil {
		t = NewTrie()
	}

	if len(word) == index {
		t.word = word
		t.count += 1

	} else {

		character := word[index]
		t.children[character] = t.children[character].insertHelper(word, index + 1)

	}

	return t
}

func (t *Trie) Get(word string) *Trie {
	return t.get(word, 0)
}

func (t *Trie) get(word string, index int) *Trie {

	if t == nil {
		return nil

	} else if index == len(word) {
		return t
	}

	character := word[index]
	return t.children[character].get(word, index + 1)
}

func (t *Trie) Contains(word string) bool {
	return t.contains(word, 0)
}

func (t *Trie) contains(word string, index int) bool {

	if t == nil {

		return false

	} else if index == len(word) {

		return true
	}

	character := word[index]
	return t.children[character].contains(word, index + 1)
}

func (t *Trie) AllKeys() []string {
	return t.KeysWithPrefix("")
}

func (t *Trie) KeysWithPrefix(prefix string) []string {

	root := t.Get(prefix)
	results := root.collect(prefix)

	sort.Sort(ByDescendingCount(results))

	words := make([]string, 0)
	for _, node := range results {
		words = append(words, node.word)
	}
	return words
}

func (t *Trie) collect(prefix string) []*Trie {

	results := make([]*Trie, 0)

	if t == nil {

		return results

	} else if t.word != "" {

		results = append(results, t)
	}

	for _, child := range t.children {

		for _, trie := range child.collect(prefix) {

			results = append(results, trie)
		}
	}

	return results
}
