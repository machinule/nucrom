package net

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
	"golang.org/x/net/context"
	"sync"
)

// gameServer contains the state of the game, and implements the GameService.
type gameServer struct {
	settings *pb.GameSettings

	lock    sync.RWMutex
	state   *pb.GameState
	turn    *pb.TurnState
	players <-chan pb.Player
}

func New(s *pb.GameSettings) (*gameServer, error) {
	p := make(chan pb.Player)
	p <- pb.Player_USA
	p <- pb.Player_USSR
	return &gameServer{
		settings: s,
		state:    nil,
		turn: &pb.TurnState{
			Index: 0,
			Moved: nil,
		},
		players: p,
	}, nil
}

func (s *gameServer) JoinGame(ctx context.Context, req *pb.JoinGameRequest) (*pb.JoinGameResponse, error) {
	p, ok := <-s.players
	if !ok {
		return nil, fmt.Errorf("No more player slots available.")
	}
	return &pb.JoinGameResponse{
		Player:   p,
		Settings: s.settings,
	}, nil
}

func (s *gameServer) GetGameState(ctx context.Context, req *pb.GetGameStateRequest) (*pb.GetGameStateResponse, error) {
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

func (s *gameServer) SetTurn(ctx context.Context, req *pb.SetTurnRequest) (*pb.SetTurnResponse, error) {
	return &pb.SetTurnResponse{}, nil
}
