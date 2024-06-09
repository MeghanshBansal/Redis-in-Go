package main

import (
	"errors"
	"log"
	"net"
)

type Peer struct {
	Conn  net.Conn
	Name  string
	msgCh chan Message
}

func NewPeer(conn net.Conn, Name string, msgCh chan Message) *Peer {
	return &Peer{
		Conn:  conn,
		Name:  Name,
		msgCh: msgCh,
	}
}

func (p *Peer) ReadFromPeer() error {
	buf := make([]byte, 1024)
	for {
		n, err := p.Conn.Read(buf)
		if err != nil {
			log.Println("failed to read message from the peer: ", p.Name)
			return errors.New("failed to read")
		}

		msgBuf := make([]byte, n)
		copy(msgBuf, buf[:n])
		p.msgCh <- Message{Peer: p.Name, Msg: string(msgBuf)}

	}
}
