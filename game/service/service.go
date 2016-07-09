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
  moves   map[pb.Player][]*pb.GameMove
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
    moves: make(map[pb.Player][]*pb.GameMove),
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

func (s *Service) advanceTurn() {
  s.lock.Lock()
  defer s.lock.Unlock()
  if len(s.turn.Moved) != 2 {
    return
  }
  s.turn = &pb.TurnState {
    Index: s.turn.Index + 1,
    Moved: nil,
  }
  s.moves = make(map[pb.Player][]*pb.GameMove)
}

func (s *Service) SubmitTurn(ctx context.Context, req *pb.SubmitTurnRequest) (*pb.SubmitTurnResponse, error) {
  s.lock.Lock()
  defer s.lock.Unlock()
  if got, want := req.TurnIndex, s.turn.Index; got != want {
    return nil, fmt.Errorf("Cannot submit a turn that is not the current turn. Got %d, want %d.", got, want)
  }
  for _, p := range s.turn.Moved {
    if p == req.Player {
      return nil, fmt.Errorf("Turn has already been submitted.")
    }
  }
  s.moves[req.Player] = req.Move
  s.turn.Moved = append(s.turn.Moved, req.Player)
  if len(s.turn.Moved) == 2 {
    go s.advanceTurn()
  }
	return &pb.SubmitTurnResponse{}, nil
}

func (s *Service) CancelTurn(ctx context.Context, req *pb.CancelTurnRequest) (*pb.CancelTurnResponse, error) {
  s.lock.Lock()
  defer s.lock.Unlock()
  if got, want := req.TurnIndex, s.turn.Index; got != want {
    return nil, fmt.Errorf("Cannot cancel a turn that is not the current turn. Got %d, want %d.", got, want)
  }
  index := -1
  for i, p := range s.turn.Moved {
    if p == req.Player {
      index = i
      break
    }
  }
  if index == -1 {
    return nil, fmt.Errorf("Cannot cancel a turn that has not been submitted.")
  }
  s.turn.Moved = append(s.turn.Moved[:index], s.turn.Moved[index + 1:]...)
  delete(s.moves, req.Player)
	return &pb.CancelTurnResponse{}, nil
}