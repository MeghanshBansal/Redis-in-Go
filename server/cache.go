package main

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	m  map[string]string
	mu *sync.RWMutex
}

type CacheI interface {
	HandleSetCommand(cmd Command)
	HandleGetCommand(cmd Command) (string, error)
	HandleDelCommand(cmd Command)
	HandleHasCommand(cmd Command) bool
}

func NewCache() CacheI {
	return Cache{
		m:  make(map[string]string),
		mu: &sync.RWMutex{},
	}
}

func (c Cache) HandleSetCommand(cmd Command) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[cmd.key] = cmd.value
	go c.deleteKeyAfter(cmd.key, cmd.t)
}

func (c Cache) HandleGetCommand(cmd Command) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if val, ok := c.m[cmd.key]; ok {
		return val, nil
	}
	return "", errors.New("key not found")
}

func (c Cache) HandleDelCommand(cmd Command) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.m, cmd.key)
}

func (c Cache) HandleHasCommand(cmd Command) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.m[cmd.key]; ok {
		return true
	}
	return false
}

func (c Cache) deleteKeyAfter(key string, t time.Duration) {
	<-time.After(t)
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}
