package main

import (
	"time"
)

const (
	CommandSet = "SET"
	CommandGet = "GET"
	CommandDel = "DEL"
	CommandHas = "HAS"
)

type Command struct {
	CMD   string
	t     time.Duration
	key   string
	value string
}