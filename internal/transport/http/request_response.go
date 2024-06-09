package http

type MessageRequest struct {
	To        string `json:"to,omitempty"`
	Message   string `json:"message"`
	Broadcast bool   `json:"broadcast,omitempty"`
}
