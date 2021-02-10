package node

import (
	"log"
	"os/exec"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) PrintDir(ctx context.Context, message *GetVersionMessage) (*GetVersionResponse, error) {
	out, err := exec.Command("cardano-cli", "--version").Output()

	if err != nil {
		log.Fatal(err)
	}

	return &GetVersionResponse{Body: string(out)}, nil
}
