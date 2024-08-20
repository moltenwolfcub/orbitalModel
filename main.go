package main

import (
	"log/slog"

	"github.com/moltenwolfcub/orbitalModel/game"
)

func main() {
	instance := game.NewGame()

	if err := instance.Run(); err != nil {
		slog.Error("Failed to execute tick: " + err.Error())
		return
	}
}
