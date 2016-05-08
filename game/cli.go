// A command line interface to play the game.
package main

import (
	"fmt"
	"github.com/turret-io/go-menu/menu"
)

func netHost(_ ...string) error {
	fmt.Println("net host")
	return nil
}

func netConnect(_ ...string) error {
	fmt.Println("net connect")
	return nil
}

func hotseat(_ ...string) error {
	fmt.Println("hotseat")
	return nil
}

func main() {
	fmt.Println("Nuclear Romance")
	fmt.Println("---------------")

	opts := []menu.CommandOption{
		menu.CommandOption{"nethost", "Host a net game.", netHost},
		menu.CommandOption{"netconnect", "Connect to a net game.", netConnect},
		menu.CommandOption{"hotseat", "Run a hotseat game.", hotseat},
	}

	menu.NewMenu(opts, menu.NewMenuOptions("Game Choice > ", 0)).Start()
}
