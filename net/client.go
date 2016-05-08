package net

import (
	pb "github.com/machinule/nucrom/proto/gen"
	"google.golang.org/grpc"
)

type GameClient struct {
	conn     *grpc.ClientConn
	service  *pb.GameServiceClient
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
	c.service = &service
	return c, nil
}

func (c *GameClient) Close() {
	c.conn.Close()
}
