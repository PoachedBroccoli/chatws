package controller

import (
	"chat_websocket/model"
	"chat_websocket/tool"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

var userList = []string{}
var users = make(map[string][]string)

// connection : user connection
type connection struct {
	id   string
	ws   *websocket.Conn
	send chan []byte
	data *model.Data
}

func (c *connection) writer() {
	for message := range c.send {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			Hub.unregister <- c
			break
		}
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "ping":
			// TODO
		case "enter":
			// append user
			c.data.User = c.data.Content
			c.data.From = c.data.Content
			// userList = append(userList, c.data.User)
			users[c.data.Group] = append(users[c.data.Group], c.data.User)
			c.data.UserList = users[c.data.Group]
			data, _ := json.Marshal(c.data)
			Hub.broadcast <- data
		case "chat":
			// chat
			data, _ := json.Marshal(c.data)
			Hub.broadcast <- data
		case "exit":
			// del user
			c.data.Type = "exit"
			// userList = remove(userList, c.data.User)
			users[c.data.Group] = tool.Remove(users[c.data.Group], c.data.User)
			c.data.UserList = users[c.data.Group]
			c.data.Content = c.data.User
			data, _ := json.Marshal(c.data)
			Hub.broadcast <- data
			Hub.unregister <- c
		}
	}
}

// Upgrader:Upgrade http request to ws request
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 2048,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSHandler : handle connection
func WSHandler(w http.ResponseWriter, r *http.Request) {
	// TODO:middleware
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	c := &connection{
		send: make(chan []byte, 128),
		ws:   ws,
		data: &model.Data{},
	}

	// register
	Hub.register <- c
	go c.writer()
	c.reader()
	defer func() {
		// del user
		c.data.Type = "exit"
		// userList = remove(userList, c.data.User)
		users[c.data.Group] = tool.Remove(users[c.data.Group], c.data.User)
		c.data.UserList = users[c.data.Group]
		c.data.Content = c.data.User
		data, _ := json.Marshal(c.data)
		Hub.broadcast <- data
		Hub.unregister <- c
	}()
}
