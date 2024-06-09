package main 

import (
	"strconv"
	"strings"
	"time"
	"errors"
	"log"
)

type Message struct {
	Peer *Peer
	Msg  string
}

func (m *Message) parseCommand() (Command, error) {
	commandArray := strings.Split(m.Msg, " ")
	var com Command
	if len(commandArray) == 4 && commandArray[0] == CommandSet {
		com.CMD = CommandSet
		durationInSeconds, err := strconv.Atoi(commandArray[1])
		if err != nil {
			log.Println("Failed to parse time in command")
			return Command{}, errors.New("failed to parse command")
		}
		com.t = time.Duration(durationInSeconds) * time.Second
		com.key = commandArray[2]
		com.value = commandArray[3]
	} else if len(commandArray) == 2 {
		switch commandArray[0] {
		case CommadnGet:
			com = Command{CMD: CommadnGet, key: commandArray[1]}
		case CommandDel:
			com = Command{CMD: CommandDel, key: commandArray[1]}
		case CommandHas:
			com = Command{CMD: CommandHas, key: commandArray[1]}
		default:
			return Command{}, errors.New("invalid command")
		}
	} else {
		return Command{}, errors.New("invalid command")
	}
	return com, nil
}