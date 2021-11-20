package lsm

import "errors"

type Store map[string]string

func New() *Store {
	s := make(Store)

	return &s
}

func (s Store) Length() int {
	return len(s)
}

func (s Store) Set(key string, value string) (Store, error) {
	if key == "" {
		return nil, errors.New("Empty characters are not allowed in key.")
	}

	s[key] = value

	return s, nil
}

func (s Store) Get(key string) (string, error) {
	v := s[key]

	if v == "" {
		return "", errors.New("A non-existent key has been specified.")
	}

	return v, nil
}

func (s Store) Remove(key string) (Store, error) {
	l := len(s)

	delete(s, key)

	l2 := len(s)

	if l == l2 {
		return nil, errors.New("It seems that you are trying to delete a key that does not exist.")
	}

	return s, nil
}

func (s Store) Clear() (Store, error) {
	for key, _ := range s {
		s.Remove(key)
	}

	return s, nil
}
