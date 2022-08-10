package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *PingMessage) (*PingMessage, error) {
	log.Print("Recieved message from cilent: %s", message.Greeting)
	return &PingMessage{Greeting: "Helo from server !!!!!!!!"}, nil
}
