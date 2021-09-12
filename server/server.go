package server

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/subchen/go-log"
	"net/http"
	"time"
)

var WSServer *socketio.Server

// Init - Web socket server initialization
func Init() {
	InitStack()

	pt := polling.Default

	wt := websocket.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	s, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			pt,
			wt,
		},
	})

	if err != nil {
		log.Fatalf("cant start Socket.IO server: %v", err)
	}

	s.OnConnect("/", connection)
	WSServer = s

	go run()
}

// Connection handling method
func connection(c socketio.Conn) error {
	start := time.Now()
	defer log.Debugf("Processed connection from %v (%v)", c.RemoteAddr(), time.Since(start))

	hub := hubs.GetActualHub()
	client := NewClient(c.ID(), hub.ID.String())
	hub.AddClient(&client)

	WSServer.JoinRoom("/", hub.ID.String(), c)

	return nil
}

func run() {
	defer shutdown()
	if err := WSServer.Serve(); err != nil {
		log.Fatalf("cant start WSServer Socket.IO: %v", err)
	}
}

func shutdown() {
	if err := WSServer.Close(); err != nil {
		log.Fatalf("cant stop WSServer Socket.IO: %v", err)
	}
}

// SendMessage - sending a message to the hub
func SendMessage(msg string, hubID string) string {
	ok := WSServer.BroadcastToRoom("/", hubID, "test", msg)
	if !ok {
		fmt.Println("cant send message in room")
	}

	return fmt.Sprintf("message sent to %s", hubID)
}

// SendMessageToUser - sending a message to the specified client
func SendMessageToUser(cID string, msg string) string {
	client := hubs.GetClientByID(cID)

	WSServer.ForEach("/", client.HubID, func(conn socketio.Conn) {
		if client.ConnectionID == conn.ID() {
			conn.Emit("test2", msg, 1)
		}
	})

	return fmt.Sprintf("Message sended to %v", client.ID.String())
}
