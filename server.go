package main

import (
	"errors"
	"log"
	"net"

	"github.com/google/uuid"
)

const DefaultAddr = ":8080"

type Config struct {
	ListenAddr string
}

type Server struct {
	Cfg       Config
	Listener  net.Listener
	Peers     map[Peer]bool
	addPeerCh chan *Peer
	quitCh    chan *Peer
	msgCh     chan Message

	Cache CacheI
}

func NewServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = DefaultAddr
	}

	cache := NewCache()

	return &Server{
		Cfg:       cfg,
		Peers:     make(map[Peer]bool),
		addPeerCh: make(chan *Peer),
		msgCh:     make(chan Message),
		Cache:     cache,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.Cfg.ListenAddr)
	if err != nil {
		log.Printf("Failed to start the server at port %s with error %s\n", s.Cfg.ListenAddr, err)
		return errors.New("failed to start the server")
	}
	log.Println("Server Created Succesfully on port ", s.Cfg.ListenAddr)
	s.Listener = ln
	go s.AddPeer()

	return s.AcceptConnection()
}

func (s *Server) AcceptConnection() error {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %s\n", err.Error())
			return errors.New("failed to accept connection")
		}
		go s.HandleConn(&conn)
	}
}

func (s *Server) HandleConn(conn *net.Conn) error {
	peer := NewPeer(*conn, uuid.NewString(), s.msgCh)
	s.addPeerCh <- peer
	if err := peer.ReadFromPeer(); err != nil {
		return err
	}
	return nil
}

func (s *Server) handleRawMessage(rawMsg Message) error {
	cmd, err := rawMsg.parseCommand()
	if err != nil {
		log.Println("failed to parse Command from user ", rawMsg.Peer)
		return  err
	}

	switch cmd.CMD {
	case CommandSet:
		s.Cache.HandleSetCommand(cmd)
	case CommadnGet:
		_, err := s.Cache.HandleGetCommand(cmd)
		if err != nil {
			log.Println("Command not found")
		}
		// todo: something with res
	case CommandDel:
		s.Cache.HandleDelCommand(cmd)
	case CommandHas:
		_ = s.Cache.HandleHasCommand(cmd)
		// todo: something with the res
	}

	return  nil
}

func (s *Server) AddPeer() {
	for {
		select {
		case rawMsg := <-s.msgCh:
			if err := s.handleRawMessage(rawMsg); err != nil {
				log.Println("failed to decode raw message received")
			}
			log.Printf("msg from %s: %s\n", rawMsg.Peer, rawMsg.Msg)
		case peer := <-s.quitCh:
			log.Printf("%s Peer has left\n", peer.Name)
			return
		case peer := <-s.addPeerCh:
			s.Peers[*peer] = true
		}
	}
}
