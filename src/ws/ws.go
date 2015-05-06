package ws

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"model"
	"net/http"
)

var S *Server

// Chat server.
type Server struct {
	pattern string
	clients map[int]*Client
	addCh   chan *Client
	delCh   chan *Client
	errCh   chan error
}

// Create new chat server.
func NewServer(pattern string) *Server {
	clients := make(map[int]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	errCh := make(chan error)

	return &Server{
		pattern,
		clients,
		addCh,
		delCh,
		errCh,
	}
}

func (s *Server) SendProjects(msg []*model.Project) {
	for _, c := range s.clients {
		c.sendProjects(msg)
	}
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen() {

	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		client := &Client{maxId, ws, s}
		maxId++
		s.addCh <- client
		client.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))
	log.Println("Created handler")

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			log.Println("Added new client", c.id)
			s.clients[c.id] = c
			log.Println("Now", len(s.clients), "clients connected.")

		// del a client
		case c := <-s.delCh:
			log.Println("Delete client", c.id)
			delete(s.clients, c.id)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		}
	}
}

//const channelBufSize = 100

var maxId int = 0

// Chat client.
type Client struct {
	id     int
	ws     *websocket.Conn
	server *Server
}

// Listen Write and Read request via chanel
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {
			//TODO: move send projects here
		}
	}
}

func (c *Client) sendProjects(msg []*model.Project) {
	log.Println("Send", c.id, ":", msg)
	websocket.JSON.Send(c.ws, msg)
}

// Listen read request via chanel
func (c *Client) listenRead() {
	log.Println("Listening read from client")
	for {
		select {

		// read data from websocket connection
		default:
			var msg SomeMessage
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF {
				c.server.delCh <- c
				return
			} else if err != nil {
				c.server.errCh <- err
			} else {
				log.Println("Received", c.id, ":", msg)
			}
		}
	}
}

type SomeMessage struct {
	Body string `json:"body"`
}

func (self SomeMessage) String() string {
	return fmt.Sprintf("%#v", self)
}
