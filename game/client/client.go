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
}

// A Client contains the necessary state for a game client.
type client struct {
	address  string
	service  pb.GameServiceClient
	settings *pb.GameSettings
	player   pb.Player
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

func (c *client) EndTurn() error {
	return nil
}
