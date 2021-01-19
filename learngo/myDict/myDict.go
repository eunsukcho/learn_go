package myDict

import "errors"

type Dictionary map[string]string

var (
	errNotFound    = errors.New("Not Found")
	errCantUpdate  = errors.New("Cant update non-exising word")
	errWordExists  = errors.New("Thant word already exists")
	errDelNotFound = errors.New("Noting to Delete ")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound: // dict에 아직 단어가 없다면
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
	/*
		if err == errNotFound {
			d[word] = def
		} else if err == nil {
			return errWordExists
		}
		return nil
	*/
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errDelNotFound
	case nil:
		delete(d, word)
	}
	return nil
}
