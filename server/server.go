package server

import (
	"fmt"
	"net"
	"time"
)

func NewServer(Addr string) *Server {
	return &Server{
		listenAddr: Addr,
		quitch:     make(chan struct{}),
		history:    make(chan Message, 10),
		clients:    make(map[net.Conn]*Client),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	fmt.Println("Listening on the port", s.listenAddr)
	defer ln.Close()
	s.ln = ln
	go s.AcceptLoop()

	// go func() {
	// 	if err := s.BroadcastMessage(); err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }()

	<-s.quitch

	return nil
}

func (s *Server) AcceptLoop() {
	for {

		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Error accepting port", err)
			continue
		}
		go s.loopRead(conn)

	}
}

func (s *Server) loopRead(conn net.Conn) {
	// pattern := fmt.Sprintf("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + s.clients[conn].Name + "]:")
	// conn.Write([]byte(pattern))
	defer func() {
		clientName := s.getClientName(conn)
		s.removeClient(conn)
		s.notifyUserLeft(clientName)

		conn.Close()
	}()

	buf := make([]byte, 2048)
	clientName, err := s.AddName(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(clientName)

	s.mu.Lock()
	s.clients[conn] = &Client{Name: clientName, Conn: conn}
	s.mu.Unlock()
	pattern := fmt.Sprintf("[%s][%s]:\n", time.Now().Format("2006-01-02 15:04:05"), clientName)
	conn.Write([]byte(pattern))

	// s.notifyNewUser(clientName)
	// pattern := fmt.Sprintf("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + clientName + "]:")
	// conn.Write([]byte(pattern))

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("User has left chat")
			return
		}
		if n == 0 {
			continue
		}

		s.history <- Message{
			Type:     "",
			From:     clientName,
			ConnTime: time.Now(),
			Payload:  buf[:n],
		}

		go func() {
			pattern := fmt.Sprintf("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + s.clients[conn].Name + "]:\n")
			conn.Write([]byte(pattern))
			if err := s.BroadcastMessage(); err != nil {
				fmt.Println(err)
				return
			}
		}()

	}
}
