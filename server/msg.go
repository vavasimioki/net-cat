package server

import (
	"fmt"
	"net"
	"os"
)

func (s *Server) notifyUserJoined(Client string) {
	notification := fmt.Sprintf("%s has been joined the chat...\n", Client)
	for _, client := range s.clients {
		client.Conn.Write([]byte(notification))
	}
}

func (s *Server) notifyUserLeft(Client string) {
	notification := fmt.Sprintf("%s has been removed from the chat...\n", Client)

	for _, client := range s.clients {
		client.Conn.Write([]byte(notification))
	}
}

func (s *Server) getClientName(conn net.Conn) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	client, ok := s.clients[conn]
	if ok {
		return client.Name
	}

	return ""
}

func (s *Server) removeClient(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, conn)
}

func (s *Server) BroadcastMessage() error {
	for {
		msg, ok := <-s.history
		if !ok {
			fmt.Println("error listening channel")
		}

		s.mu.Lock()
		for conn, client := range s.clients {
			if client.Name != msg.From {
				conn.Write([]byte(s.sendMessage(msg)))
			}
		}
		s.mu.Unlock()

	}
}

func (s *Server) sendMessage(msg Message) string {
	return fmt.Sprintf("["+msg.ConnTime.Format("2006-01-02 15:04:05")+"]"+"["+msg.From+"]:"+"%s", msg.Payload)
}

func LoadLogo() (string, error) {
	logo, err := os.ReadFile("logo.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	return string(logo), nil
}
