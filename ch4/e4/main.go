package main

import (
	"fmt"
)
type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (auth authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", auth.username, auth.password)
}

func main() {
	me := authenticationInfo {
		username: "muhammed",
		password: "1234",
	}
	fmt.Println(me.getBasicAuth())
}
