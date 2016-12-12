package main

import "fmt"

type Node struct {
	wordEndsHere bool
	children     map[byte]*Node
}

func NewNode() *Node {

	children := make(map[byte]*Node)

	return &Node{children:children}
}

func (t *Node) Insert(word string, index int) {

	letter := word[index]

	if _, ok := t.children[letter]; !ok {
		t.children[letter] = NewNode()
	}

	if index + 1 == len(word) {

		t.wordEndsHere = true

	} else {
		t.children[letter].Insert(word, index+1)
	}
}

func (t *Node) Contains(word string, index int) bool {

	if len(word) == 0 {
		return true
	}

	if len(word)-1 == index {
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
		return t.children[letter].Contains(word, index+1)
	}

}



func main() {

	root := NewNode()
	root.Insert("hello", 0)
	root.Insert("how are you", 0)
	root.Insert("i am good", 0)
	root.Insert("thanks", 0)

	fmt.Println(root.Contains("thanks", 0))
}