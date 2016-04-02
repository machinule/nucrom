package main

import "github.com/machinule/nucrom/frontend/sdl2"

func main() {
	fe, err := sdl2.New()
	if err != nil {
		panic(err)
	}
	if err := fe.Run(); err != nil {
		panic(err)
	}
}
