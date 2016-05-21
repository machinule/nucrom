// A command line interface to play the game.
package main

import (
	"fmt"
	"log"

	gameclient "github.com/machinule/nucrom/game/client"
	gameserver "github.com/machinule/nucrom/game/server"
	"github.com/machinule/nucrom/game/setup"
	"github.com/turret-io/go-menu/menu"
)

type game struct {
	server *gameserver.Server
	client gameclient.Client
}

func (g *game) StartServerOrDie(port string) {
	g.server = gameserver.New(port, setup.CreateGameSettings())
	if err := g.server.Start(); err != nil {
		log.Fatalf("Starting game server: %v", err)
	}
}

func (g *game) netHost(_ ...string) error {
	fmt.Println("Hosting a net game...")
	g.server = gameserver.New(":7544", setup.CreateGameSettings())
	g.client = gameclient.New("localhost:7544")
	g.GameOn()
	return nil
}

func (g *game) netConnect(_ ...string) error {
	fmt.Println("Connecting to a net game...")
	fmt.Print("Hostport: ")
	var hostport string
	fmt.Sscanln("%s", &hostport)
	g.client = gameclient.New(hostport)
	g.GameOn()
	return nil
}

func (g *game) hotseat(_ ...string) error {
	fmt.Println("Starting a hotseat game...")
	g.server = gameserver.New(":7544", setup.CreateGameSettings())
	g.client = gameclient.NewMux("localhost:7544", 2)
	g.GameOn()
	return nil
}

func (g *game) GameOn() {
	if g.server != nil {
		if err := g.server.Start(); err != nil {
			log.Fatalf("Failed to start game server: %v", err)
		}
	}
	if err := g.client.Connect(); err != nil {
		log.Fatalf("Failed to connect to game server: %v", err)
	}
	if err := g.client.Join(); err != nil {
		log.Fatalf("Failed to join game: %v", err)
	}
	fmt.Println("gaming on...")
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
}
