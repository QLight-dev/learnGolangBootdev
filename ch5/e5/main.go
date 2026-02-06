package main

import "fmt"

type formatter interface {
	format() string
}

type plainText struct {
	message string
}
type bold struct {
	message string
}
type code struct {
	message string
}

func (msg plainText) format() string {
	return msg.message
}
func (msg bold) format() string {
	return fmt.Sprintf("**%s**", msg.message)
}
func (msg code) format() string {
	return fmt.Sprintf("`%s`", msg.message)
}

// Don't Touch below this line

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}

