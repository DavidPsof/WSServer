package server

import (
	"WSServer/config"
	"WSServer/domain"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"github.com/subchen/go-log"
	"net/http"
	"strconv"
)

// ServerManager - describe server manager struct
type ServerManager struct {
	Connections map[string]*websocket.Conn
}

// NewServerManager - create new server manager
func NewServerManager() *ServerManager {
	manager := ServerManager{}
	manager.Connections = make(map[string]*websocket.Conn, 0)

	return &manager
}

// SocketConnection - handle request on creating new ws connection
func (s *ServerManager) SocketConnection(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}

	t := r.URL.Query().Get("token")

	if t[:6] != "bearer" {
		return
	}

	tokenString := t[7:]

	token, _ := jwt.ParseWithClaims(tokenString, &domain.JwtInfo{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неизвестный метод подписи: %v", token.Header["alg"])
		}

		return []byte(config.Get().JwtKey), nil
	})

	custromClaims := token.Claims.(*domain.JwtInfo)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Error during connection upgradation: %v", err)
		return
	}

	err = conn.WriteMessage(1, []byte("pong"))
	if err != nil {
		log.Errorf("Error of sending message: %v", err)
		return
	}

	s.Connections[strconv.Itoa(custromClaims.UserID)] = conn
	log.Debugf("Connection with token - %v created", strconv.Itoa(custromClaims.UserID))
}
