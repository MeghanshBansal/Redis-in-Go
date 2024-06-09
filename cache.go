package main

import "sync"

type Cache struct {
	c  map[string]string
	mu *sync.RWMutex
}

type CacheI interface {
	HandleSetCommand(cmd Command) (bool, error)
	HandleGetCommand(cmd Command) (string, error)
	HandleDelCommand(cmd Command) (bool, error)
	HandleHasCommand(cmd Command) (bool, error)
}

func NewCache() CacheI {
	return Cache{
		c:  make(map[string]string),
		mu: &sync.RWMutex{},
	}
}

func (c Cache) HandleSetCommand(cmd Command) (bool, error) {
	return true, nil
}
func (s Cache) HandleGetCommand(cmd Command) (string, error) {
	return "", nil
}
func (s Cache) HandleDelCommand(cmd Command) (bool, error) {
	return true, nil
}
func (s Cache) HandleHasCommand(cmd Command) (bool, error) {
	return true, nil
}
