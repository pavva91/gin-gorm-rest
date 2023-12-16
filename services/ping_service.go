package services

import (
	"fmt"
)

const (
	pong = "pong"
)

var (
	Ping pingServicer = ping{}
)

type pingServicer interface {
	HandlePing() (string, error)
}

type ping struct{}

func (s ping) HandlePing() (string, error) {
	fmt.Println("doing some complex things...")
	return pong, nil
}
