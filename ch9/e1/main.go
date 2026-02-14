package main

import (
	"errors"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)

	if len(names) != len(phoneNumbers) {
		return userMap, errors.New("invalid sizes")
	}
	for i, name := range names {
		userMap[name] = user{name: name, phoneNumber: phoneNumbers[i]}
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
