package net

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GameClient struct {
	conn     *grpc.ClientConn
	service  pb.GameServiceClient
	settings *pb.GameSettings
	player   pb.Player
}

func NewClient(address string) (*GameClient, error) {
	c := &GameClient{}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c.conn = conn
	service := pb.NewGameServiceClient(c.conn)
	c.service = service
	return c, nil
}

func (c *GameClient) Join() error {
	r, err := c.service.JoinGame(context.Background(), &pb.JoinGameRequest{})
	if err != nil {
		return err
	}
	c.settings = r.Settings
	c.player = r.Player
	fmt.Printf("Playing as: %v\n", c.player)
	return nil
}

func (c *GameClient) Close() {
	c.conn.Close()
}
