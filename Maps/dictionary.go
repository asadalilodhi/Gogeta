package main

import "errors"

type Dictionary map[string]string


var ErrNotFound = errors.New("could not find the word you were looking for")

func (dictionary Dictionary) search(word string) (string, error) {
	word, found := dictionary[word]
	if !found {
		return "", ErrNotFound
	}

	return word, nil
}