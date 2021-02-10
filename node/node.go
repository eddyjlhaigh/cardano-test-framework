package node

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) PrintDir(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received: %s", message.Body)
	return &Message{Body: "Hello!"}, nil
}
