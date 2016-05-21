// Package client implements a game client.
package client

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client interface {
	Connect() error
	Join() error
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
	s, err := c.service.GetGameState(context.Background(), &pb.GetGameStateRequest{
		ReturnTurnOnly: false,
	})
	if err != nil {
		return fmt.Errorf("getting initial game state: %v", err)
	}
	c.turn = s.Turn
	c.state = s.State
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

func (c *client) EndTurn() error {
	c.turn.Index = c.turn.Index + 1
	return nil
}

func (c *client) GameOver() bool {
	return c.Turn() > 40
}
