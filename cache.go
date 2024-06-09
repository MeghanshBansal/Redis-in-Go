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
	HandleSetCommand(cmd Command) (bool, error)
	HandleGetCommand(cmd Command) (string, error)
	HandleDelCommand(cmd Command) (bool, error)
	HandleHasCommand(cmd Command) (bool, error)
}

func NewCache() CacheI {
	return Cache{
		m:  make(map[string]string),
		mu: &sync.RWMutex{},
	}
}

func (c Cache) HandleSetCommand(cmd Command) (bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[cmd.key] = cmd.value
	go c.deleteKey(cmd.key, cmd.t)
	return true, nil
}


func (c Cache) HandleGetCommand(cmd Command) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if val, ok := c.m[cmd.key]; ok {
		return val, nil
	}
	return "", errors.New("key not found")
}


func (c Cache) HandleDelCommand(cmd Command) (bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.m, cmd.key)
	return true, nil
}

func (c Cache) HandleHasCommand(cmd Command) (bool, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.m[cmd.key]; ok {
		return true, nil
	}
	return false, nil
}

func (c Cache) deleteKey(key string, t time.Duration) {
	<-time.After(t)
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}
