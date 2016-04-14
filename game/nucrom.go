package nucrom

import (
    "log"

    "github.com/machinule/nucrom/frontend/sdl2"

    "github.com/machinule/nucrom/game/setup"
)

func ui() {
	fe, err := sdl2.New()
	if err != nil {
		panic(err)
	}
	if err := fe.Run(); err != nil {
		panic(err)
	}
}

func Run() {
    setup.Try()

    gameSettings := setup.CreateGameSettings()
    log.Printf("Game settings: ", gameSettings)
}
