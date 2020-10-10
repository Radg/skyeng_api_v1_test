package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getWords(word string) ([]Word, error) {
	resp, err := http.Get(words_url + "?search=" + word)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var words []Word
	err = json.Unmarshal(body, &words)
	if err != nil {
		return nil, err
	}

	return words, nil
}

func getPartOfSpeechByCode(code string) string {
	return PartsOfSpeech[code]
}
