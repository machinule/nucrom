// Package client implements a game client.
package client

import (
	"fmt"
  "time"
	pb "github.com/machinule/nucrom/proto/gen"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client interface {
	Connect() error
	Join() error
  StartTurn() error
	EndTurn() error
	Player() pb.Player
	Turn() int
	GameOver() bool
}

// A Client contains the necessary state for a game client.
type client struct {
	address  string
	service  pb.GameServiceClient
	settings *pb.GameSettings
	player   pb.Player
	turn     *pb.TurnState
	state    *pb.GameState
}

// New creates a new Client.
func New(address string) Client {
	return &client{
		address: address,
	}
}

func (c *client) Connect() error {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("connecting new client to address '%s': %v", c.address, err)
	}
	c.service = pb.NewGameServiceClient(conn)
	return nil
}

func (c *client) Join() error {
	r, err := c.service.JoinGame(context.Background(), &pb.JoinGameRequest{})
	if err != nil {
		return fmt.Errorf("joining game: %v", err)
	}
	c.settings = r.Settings
	c.player = r.Player
	return nil
}

func (c *client) Turn() int {
	if c.turn != nil {
		return int(c.turn.Index)
	}
	return 0
}

func (c *client) Player() pb.Player {
	return c.player
}

func (c *client) StartTurn() error {
  for {
    s, err := c.service.GetGameState(context.Background(), &pb.GetGameStateRequest{
      ReturnTurnOnly: true,
    })
    if err != nil {
      return fmt.Errorf("getting latest turn from server: %v", err)
    }
    if c.turn == nil || s.Turn.Index != c.turn.Index {
      s, err = c.service.GetGameState(context.Background(), &pb.GetGameStateRequest{
        ReturnTurnOnly: false,
      })
      if err != nil {
        return fmt.Errorf("getting game state from server: %v", err)
      }
      c.turn = s.Turn
      c.state = s.State
      break
    }    
    time.Sleep(1 * time.Second)
  } 
  return nil
}

func (c *client) EndTurn() error {
  _, err := c.service.SubmitTurn(context.Background(), &pb.SubmitTurnRequest{
    Player: c.player,
    TurnIndex: c.turn.Index,
  })
  if err != nil {
    return fmt.Errorf("ending turn: %v", err)
  }
	return nil
}

func (c *client) GameOver() bool {
	return c.Turn() > 10
}
