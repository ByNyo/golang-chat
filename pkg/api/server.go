package api

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"sync"
)

type Server interface {
	HandleWS(ws *websocket.Conn)
}

type server struct {
	conns map[*websocket.Conn]bool
	mu    sync.Mutex
}

func NewServer() Server {
	return &server{conns: make(map[*websocket.Conn]bool)}
}

func (s *server) HandleWS(ws *websocket.Conn) {
	fmt.Println("incoming connection from client:", ws.RemoteAddr())
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *server) broadcast(b []byte) {
	for ws := range s.conns {
		if s.conns[ws] {
			go func(ws *websocket.Conn) {
				if _, err := ws.Write(b); err != nil {
					fmt.Println("error writing to websocket:", err)
				}
			}(ws)
		}
	}
}

func (s *server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				delete(s.conns, ws)
				break
			}
			fmt.Println("read error:", err)
			continue
		}
		s.broadcast(buf[:n])
	}
}
