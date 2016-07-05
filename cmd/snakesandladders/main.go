package main

import (
	"flag"
	"github.com/scottjbarr/snakesandladders"
	"os"
)

func main() {
	player1 := flag.String("p1", "", "player 1 name")
	player2 := flag.String("p2", "", "player 2 name")
	flag.Parse()

	if *player1 == "" || *player2 == "" {
		flag.Usage()
		os.Exit(1)
	}
	game := snakesandladders.New(player1, player2)
	game.Autoplay()
}
