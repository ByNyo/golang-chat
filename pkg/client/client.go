package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/websocket"
	"net/url"
	"os"
	"strings"
	"time"
)

var cmdPrefix = "%"

func main() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Username: ")
	username, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	username = strings.TrimSpace(username)
	ws, err := Connect(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	ListenForMessages(ws, buf, username)
}

func ListenForMessages(ws *websocket.Conn, buf *bufio.Reader, username string) {
	defer ws.Close()
	go func() {
		fmt.Print("> ")
		for {
			message, err := buf.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}
			if message == cmdPrefix+"exit\n" {
				os.Exit(0)
			}
			msg := fmt.Sprintf("\r%s: %s", username, string(message))
			ws.Write([]byte(msg))
			time.Sleep(100 * time.Millisecond)
		}
	}()
	for {
		var msg = make([]byte, 1024)
		n, err := ws.Read(msg)
		if err != nil {
			return
		}
		message := string(msg[:n])
		fmt.Print(message)
		fmt.Print("> ")
	}
}

func Connect(username string) (conn *websocket.Conn, err error) {
	URL, err := url.Parse("ws://localhost:8080/ws")
	if err != nil {
		return nil, err
	}
	query := URL.Query()
	query.Add("username", username)
	URL.RawQuery = query.Encode()
	ws, err := websocket.Dial(URL.Redacted(), "", "http://localhost/")
	if err != nil {
		return nil, err
	}
	return ws, nil
}
