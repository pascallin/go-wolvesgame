package tcpsocket

import (
	"bufio"
	"fmt"
	"net"
)

type Hub struct {
	clients map[*Client]bool
	broadcast chan []byte
	register chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}


func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			fmt.Println("hub register")
			h.clients[client] = true
			go client.writePump()
			go client.readPump()
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			fmt.Println("hub broadcast")
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

type Client struct {
	hub *Hub
	conn net.Conn
	send chan []byte
}

func (c *Client) writePump() {
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// TODO
				return
			}
			Send(c.conn, string(message))
		}
	}
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		message, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			// TODO: error handler
			break
		}
		c.hub.broadcast <- []byte(message)
	}
}