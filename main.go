package main

import "fmt"

func main() {
	//var words []Word
	query := "dog"
	words, err := getWords(query)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	for _, word := range words {
		if word.Text == query {
			for j, meaning := range word.Meanings {
				fmt.Println(j, meaning.Translation.Text)
			}
		}
	}

}
