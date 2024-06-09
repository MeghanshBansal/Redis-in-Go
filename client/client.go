package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Config struct {
	ServerAddr string
}

type Client struct {
	Cfg    Config
	Conn   net.Conn
	Cmd    CommandI
	RespCh chan string
}

func NewClient(cfg Config) *Client {
	command := NewCommand()
	return &Client{
		Cfg:    cfg,
		Cmd:    command,
		RespCh: make(chan string, 0),
	}
}

func (c *Client) Start() error {
	dialer := net.Dialer{
		Timeout: 5 * time.Second,
	}

	conn, err := dialer.Dial("tcp", "localhost"+c.Cfg.ServerAddr)
	if err != nil {
		log.Printf("Failed to connect to %s: %v\n", c.Cfg.ServerAddr, err)
		return err
	}
	defer conn.Close()

	c.Conn = conn

	go c.readFromServer()
	go c.logMessage()

	return c.startCommunication()
}

func (c *Client) startCommunication() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("Enter the command to send")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read input")
			continue
		}
		cmd = strings.TrimSpace(cmd)
		cm, err := c.Cmd.ParseCommand(cmd)
		if err != nil {
			log.Println("Invalid Command")
			continue
		}

		err = c.handleCommands(cm)
		if err != nil {
			log.Println("failed to execute command")
			continue
		}
	}
}

func (c *Client) handleCommands(cm Command) error {
	var msg string
	switch cm.Cmd {
	case CommandSet:
		msg = fmt.Sprintf("%s %d %s %s\n", cm.Cmd, cm.T, cm.Key, cm.Value)

	case CommandGet, CommandDel, CommandHas:
		msg = fmt.Sprintf("%s %s", cm.Cmd, cm.Key)
	}
	_, err := c.Conn.Write([]byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) readFromServer() {
	buf := make([]byte, 2048)
	for {
		n, err := c.Conn.Read(buf)
		if err != nil {
			log.Println("failed to read from server")
			continue
		}

		msg := string(buf[:n])
		c.RespCh <- msg
	}
}

func (c *Client) logMessage() {
	for {
		select {
		case msg := <-c.RespCh:
			log.Println(msg)
		}
	}
}
