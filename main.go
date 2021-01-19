package main

import (
	"fmt"
	"learngo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}

	// Search
	/*
		definition, err := dictionary.Search("first")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	*/

	// Add
	/*
		word := "hello"
		definition := "Greeting"
		err := dictionary.Add(word, definition)
		if err != nil {
			fmt.Println(err)
		}
		h_definition, _ := dictionary.Search("hello")
		fmt.Println(h_definition)

		word2 := "hello"
		definition2 := "Greeting"
		err2 := dictionary.Add(word2, definition2)
		if err2 != nil {
			fmt.Println(err2)
		}
		h_definition2, _ := dictionary.Search("hello")
		fmt.Println(h_definition2)
	*/

	// Update
	/*
		baseWord := "hello"
		add_err := dictionary.Add(baseWord, "First")
		if add_err != nil {
			fmt.Println(add_err)
		}
		word, _ := dictionary.Search(baseWord)
		fmt.Println("before update", word)

		up_err := dictionary.Update("hello2", "Second")
		if up_err != nil {
			fmt.Println(up_err)
		}
		word2, _ := dictionary.Search(baseWord)
		fmt.Println("after update", word2)
	*/

	// Delete
	add_Word := "new Word"
	add_err := dictionary.Add(add_Word, "newContent")
	if add_err != nil {
		fmt.Println(add_err)
	}
	searchVal, _ := dictionary.Search(add_Word)
	fmt.Println("Search Value", searchVal)

	del_err := dictionary.Delete("add_Word")
	if del_err != nil {
		fmt.Println(del_err)
	}

	searchVal2, searchErr := dictionary.Search(add_Word)
	if searchErr != nil {
		fmt.Println(searchErr)
	} else {
		fmt.Println("Search Value", searchVal2)
	}
}
