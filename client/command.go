package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
)
const (
	CommandSet = "SET"
	CommandGet = "GET"
	CommandDel = "DEL"
	CommandHas = "HAS"
)

type Command struct {
	Cmd   string
	T     int64
	Key   string
	Value string
}

type CommandI interface{
	ParseCommand(cmd string) (Command, error)
}

func NewCommand() CommandI{
	return Command{
	}
}


func (c Command) ParseCommand(cmd string) (Command, error) {
	cm := strings.Split(cmd, " ")
	if len(cm) == 0 {
		return Command{}, errors.New("invalid command")
	}
	switch cm[0] {
	case CommandSet:
		if len(cm) < 3{
			return Command{}, errors.New("invalid SET command")
		}else if len(cm) ==3{
			return Command{
				Cmd: cm[0],
				T: 10,
				Key: cm[1],
				Value: cm[2],				
			}, nil
		}else {
			timeValueInSec, err := strconv.Atoi(cm[1])
			if err!=nil{
				log.Println("failed to parse the time in SET command")
				return Command{}, errors.New("failed to parse the time in SET command")
			}
			return Command{
				Cmd: cm[0],
				T: int64(timeValueInSec),
				Key: cm[2],
				Value: cm[3],				
			}, nil
		}
	case CommandGet:
		if len(cm) < 2{
			return Command{}, errors.New("invalid GET command")
		}else{
			return Command{
				Cmd: cm[0],
				Key: cm[1],
			}, nil
		}
	case CommandDel:
		if len(cm) < 2{
			return Command{}, errors.New("invalid DEL command")
		}else{
			return Command{
				Cmd: cm[0],
				Key: cm[1],
			}, nil
		}
	case CommandHas:
		if len(cm) < 2{
			return Command{}, errors.New("invalid HAS command")
		}else{
			return Command{
				Cmd: cm[0],
				Key: cm[1],
			}, nil
		}
	}

	return Command{}, errors.New("invalid command")
}