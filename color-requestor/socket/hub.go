package socket

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	clients     map[*Client]bool
	sendMessage chan []byte
	add         chan *Client
	remove      chan *Client
}

func NewHub() *Hub {
	return &Hub{
		sendMessage: make(chan []byte),
		add:         make(chan *Client),
		remove:      make(chan *Client),
		clients:     make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.add:
			h.clients[client] = true
		case client := <-h.remove:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.sendMessage:
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
