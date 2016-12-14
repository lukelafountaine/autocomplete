package main

type Node struct {
	word     string
	children map[byte]*Node
}

func NewNode() *Node {
	children := make(map[byte]*Node)
	return &Node{children:children}
}

func (t *Node) Insert(word string) {
	t.insertHelper(word, 0)
}

func (t *Node) insertHelper(word string, index int) *Node {

	if t == nil {
		t = NewNode()
	}

	if len(word) == index {
		t.word = word
		return t
	}

	character := word[index]
	t.children[character] = t.children[character].insertHelper(word, index + 1)

	return t
}

func (t *Node) Get(word string) *Node {
	return t.get(word, 0)
}

func (t *Node) get(word string, index int) *Node {

	if t == nil {
		return nil

	} else if index == len(word) {
		return t
	}

	character := word[index]
	return t.children[character].get(word, index + 1)
}

func (t *Node) Contains(word string) bool {
	return t.containsHelper(word, 0)
}

func (t *Node) containsHelper(word string, index int) bool {

	if t == nil {

		return false

	} else if index == len(word) {

		return true
	}

	character := word[index]
	return t.children[character].containsHelper(word, index + 1)
}

func (t *Node) AllKeys() []string {
	return t.KeysWithPrefix("")
}

func (t *Node) KeysWithPrefix(prefix string) []string {

	root := t.Get(prefix)
	results := root.collect(prefix)
	return results
}

func (t *Node) collect(prefix string) []string {

	results := make([]string, 0)

	if t == nil {

		return results

	} else if t.word != "" {

		results = append(results, t.word)
	}

	for _, child := range t.children {

		for _, word := range child.collect(prefix) {

			results = append(results, word)
		}
	}

	return results
}
