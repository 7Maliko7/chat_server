package tcp

type MessageRequest struct {
	To        string `json:"to,omitempty"`
	Message   string `json:"message"`
	Broadcast bool   `json:"broadcast,omitempty"`
}

type ConnectRequest struct {
	Name string `json:"name"`
}
