package server

import (
	"net"
	"sync"
	"time"
)

type Message struct {
	ConnTime time.Time
	From     string
	Payload  []byte
	Type     string
}

type Client struct {
	Name string
	Conn net.Conn
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	history    chan Message
	clients    map[net.Conn]*Client
	mu         sync.Mutex
}
