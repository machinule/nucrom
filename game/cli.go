// A command line interface to play the game.
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/machinule/nucrom/game/setup"
	gamenet "github.com/machinule/nucrom/net"
	pb "github.com/machinule/nucrom/proto/gen"
	"github.com/turret-io/go-menu/menu"
	"google.golang.org/grpc"
	"time"
)

type game struct {
	server *grpc.Server
}

func (g *game) startServer(port string) {
	fmt.Println("Starting game server...")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	g.server = grpc.NewServer()
	gameServer, _ := gamenet.NewServer(setup.CreateGameSettings())
	pb.RegisterGameServiceServer(g.server, gameServer)
	go g.server.Serve(lis)
}

func (g *game) netHost(_ ...string) error {
	fmt.Println("Hosting a net game...")
	g.startServer(":7754")
	client, _ := gamenet.NewClient("localhost:7754")
	defer client.Close()
	time.Sleep(10 * time.Second)
	return nil
}

func (g *game) netConnect(_ ...string) error {
	fmt.Println("Connecting to a net game...")
	fmt.Print("Hostport: ")
	var hostport string
	fmt.Sscanln("%s", &hostport)
	client, _ := gamenet.NewClient(hostport)
	defer client.Close()
	return nil
}

func (g *game) hotseat(_ ...string) error {
	fmt.Println("Starting a hotseat game...")
	g.startServer(":7754")
	firstClient, _ := gamenet.NewClient("localhost:7754")
	defer firstClient.Close()
	secondClient, _ := gamenet.NewClient("localhost:7754")
	defer secondClient.Close()

	return nil
}

func (g *game) Stop() {
	fmt.Println("Stopping game server...")
	g.server.Stop()
}

func main() {
	g := &game{}
	fmt.Println("Nuclear Romance")
	fmt.Println("---------------")

	opts := []menu.CommandOption{
		menu.CommandOption{"nethost", "Host a net game.", g.netHost},
		menu.CommandOption{"netconnect", "Connect to a net game.", g.netConnect},
		menu.CommandOption{"hotseat", "Run a hotseat game.", g.hotseat},
	}

	menu.NewMenu(opts, menu.NewMenuOptions("Game Choice > ", 0)).Start()
	g.Stop()
}
