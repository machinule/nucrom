// Server implements a GRPC game serving, including all services.
package server

import (
	"fmt"
	gameservice "github.com/machinule/nucrom/game/service"
	pb "github.com/machinule/nucrom/proto/gen"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	port     string
	settings *pb.GameSettings
	server   *grpc.Server
}

func New(port string, settings *pb.GameSettings) *Server {
	return &Server{
		port:     port,
		settings: settings,
	}
}

// Start starts a game server on s.port asynchronously.
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return fmt.Errorf("failed to listen on '%s': %v", s.port, err)
	}
	s.server = grpc.NewServer()
	gameService, err := gameservice.New(s.settings)
	if err != nil {
		return fmt.Errorf("failed to create game server with settings '%v': %v", s.settings, err)
	}
	pb.RegisterGameServiceServer(s.server, gameService)
	go s.server.Serve(lis)
	return nil
}

func (s *Server) Stop() {
	s.server.Stop()
}
