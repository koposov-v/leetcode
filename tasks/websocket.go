package tasks

type Message struct {
	UserId  int
	Payload string
}

type Connection struct {
	UserId   int
	DeviceId int
}

func (c *Connection) Write(p []byte) (int, error) {

	return 0, nil
}

type WSServer struct {
	connectedClientsCount uint64
	clients               map[Connection]bool
	register              chan Connection
	unregister            chan Connection
}

func (w *WSServer) Run() {
	select {
	case client := <-w.register:
		w.clients[client] = true
		break

	case client := <-w.unregister:
		if _, ok := w.clients[client]; ok {
			delete(w.clients, client)
		}
	}

}

func (w *WSServer) handleConnect(c Connection) {
	w.register <- c
}

func (w *WSServer) handleDisconnect() {

}

func (w *WSServer) handleQueueMessages(messages []Message) (int, error) {
	return 0, nil
}
