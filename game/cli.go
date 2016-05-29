// A command line interface to play the game.
package main

import (
	"fmt"
	"log"

	gameclient "github.com/machinule/nucrom/game/client"
	gameserver "github.com/machinule/nucrom/game/server"
	"github.com/machinule/nucrom/game/setup"
	"github.com/machinule/nucrom/frontend/menu"
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

func (g *game) netHost() error {
	fmt.Println("Hosting a net game...")
	g.server = gameserver.New(":7544", setup.CreateGameSettings())
	g.client = gameclient.New("localhost:7544")
	g.GameOn()
	return nil
}

func (g *game) netConnect() error {
	fmt.Println("Connecting to a net game...")
	fmt.Print("Hostport: ")
	var hostport string
	fmt.Sscanln("%s", &hostport)
	g.client = gameclient.New(hostport)
	g.GameOn()
	return nil
}

func (g *game) hotseat() error {
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
  turnMenu := menu.New([]menu.Option{
    {"placeholder", "Do nothing"},
    {"end", "End turn."},
  })
	for !g.client.GameOver() {
		fmt.Printf("\n\n\n\n\n-----\n\tPlayer: %s\n\tTurn: %d\n-----\n", g.client.Player(), g.client.Turn())
    end := false
    for !end {
      switch turnMenu.Ask() {
      case "placeholder":
        fmt.Println("doing nothing")
      default:
        end = true
      }
    } 
    g.client.EndTurn()
	}
	fmt.Println("Game Over.")
}

func main() {
	g := &game{}
	fmt.Println("Nuclear Romance")
	fmt.Println("---------------")

  m := menu.New([]menu.Option{
		{"nethost", "Host a net game."},
		{"netconnect", "Connect to a net game."},
		{"hotseat", "Run a hotseat game."},
    {"end", "End game."},
	})
  end := false
  for !end {
    switch m.Ask() {
    case "nethost":
      g.netHost()
    case "netconnect":
      g.netConnect()
    case "hotseat":
      g.hotseat()
    default:
      end = true
    }
  }
}
