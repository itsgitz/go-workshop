package main

import "fmt"

type Power struct {
	Value string
}

func main() {
	var keyArray []string

	character := make(map[string]string)
	powers := make(map[string]int)

	character["onePiece"] = "Monkey D Luffy"
	character["naruto"] = "Uzumaki Naruto"
	character["fairyTail"] = "Natsu Dragneel"
	character["unknown"] = "I don't know"

	powers["onePiece"] = 98
	powers["naruto"] = 99
	powers["fairyTail"] = 98
	powers["unknown"] = 0

	showListCharacter(character, powers)
	extractKeys(keyArray, character)

	// how to delete elements in map
	delete(character, "unknown")
	delete(powers, "unknown")

	fmt.Printf("\nAfter Deleted:\n")

	showListCharacter(character, powers)
	extractKeys(keyArray, character)
}

func showListCharacter(char map[string]string, power map[string]int) {
	for k, v := range char {
		fmt.Println(v, ":", power[k])
	}
}

func extractKeys(key []string, data map[string]string) {
	for k := range data {
		key = append(key, k)
	}

	fmt.Println("Extract keys:", key)
}
