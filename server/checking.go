package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"unicode"
)

// type server struct {
// 	server mode.Server
// }

func (s *Server) AddName(conn net.Conn) (string, error) {
	welcomeMsg, err := LoadLogo()
	if err != nil {
		return "Error reading file", err
	}
	conn.Write([]byte(welcomeMsg))

	reader := bufio.NewReader(conn)
	var name string
	for {
		name, err = reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				return "", fmt.Errorf("error reading name: %v", err)
			}
		}

		name = strings.TrimSpace(name)

		if s.CheckName(name) {
			break
		} else {
			conn.Write([]byte("[ENTER YOUR NAME]: "))
		}
	}
	s.notifyUserJoined(name)

	return name, nil
}

func (s *Server) CheckName(name string) bool {
	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
