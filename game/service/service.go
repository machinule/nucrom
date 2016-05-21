// Package service implements the GameService.
package service

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
	"golang.org/x/net/context"
	"sync"
)

// Service contains the state of the game, and implements the GameService.
type Service struct {
	settings *pb.GameSettings

	lock    sync.RWMutex
	state   *pb.GameState
	turn    *pb.TurnState
	players <-chan pb.Player
}

// New creates a new Service with the specified GameSettings.
func New(s *pb.GameSettings) *Service {
	p := make(chan pb.Player, 2)
	p <- pb.Player_USA
	p <- pb.Player_USSR
	return &Service{
		settings: s,
		state:    nil,
		turn: &pb.TurnState{
			Index: 0,
			Moved: nil,
		},
		players: p,
	}
}

func (s *Service) JoinGame(ctx context.Context, req *pb.JoinGameRequest) (*pb.JoinGameResponse, error) {
	p, ok := <-s.players
	if !ok {
		return nil, fmt.Errorf("No more player slots available.")
	}
	return &pb.JoinGameResponse{
		Player:   p,
		Settings: s.settings,
	}, nil
}

func (s *Service) GetGameState(ctx context.Context, req *pb.GetGameStateRequest) (*pb.GetGameStateResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	resp := &pb.GetGameStateResponse{
		Turn: s.turn,
	}
	if req.ReturnTurnOnly {
		return resp, nil
	}
	resp.State = s.state
	return resp, nil
}

func (s *Service) SetTurn(ctx context.Context, req *pb.SetTurnRequest) (*pb.SetTurnResponse, error) {
	return &pb.SetTurnResponse{}, nil
}
