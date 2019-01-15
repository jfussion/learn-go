package main

import "errors"

type Dictionary map[string]string

var ErrUnknownWord = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	if val, ok := d[word]; ok {
		return val, nil
	}

	return "", ErrUnknownWord
}
