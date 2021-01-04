package controller

import (
	"chat_websocket/model"
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

// hub : handle data for user
type hub struct {
	connections map[*connection]bool
	broadcast   chan []byte
	register    chan *connection
	unregister  chan *connection
}

// Hub : init
var Hub = hub{
	connections: make(map[*connection]bool),
	broadcast:   make(chan []byte),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
}

// handle data
func (h *hub) Run() {
	for {
		select {
		// register
		case c := <-h.register:
			h.connections[c] = true
			id, _ := uuid.NewV4()
			c.id = id.String()
			c.data.ID = id.String()
			c.data.IP = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = userList

			data, _ := json.Marshal(c.data)
			c.send <- data

		// unregister
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}

		// sync message
		case data := <-h.broadcast:
			tempData := &model.Data{}
			json.Unmarshal(data, tempData)
			for c := range h.connections {
				if tempData.Group == c.data.Group {
					select {
					case c.send <- data:
					default:
						delete(h.connections, c)
						close(c.send)
					}
				}
			}
		}
	}
}
