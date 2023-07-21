package maps

import "errors"

type Dictionary map[string]string

func (dictionary Dictionary) Search(keyword string) (string, error) {
	definition, ok := dictionary[keyword]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
