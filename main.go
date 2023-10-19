package main

import (
	"dadu_game/game"
	"flag"
)

func main() {

	player := flag.Int("pemain", 3, "menentuka jumlah pemain")
	dadu := flag.Int("dadu", 4, "menentukan jumlah dadu")

	flag.Parse()

	daduFactory := game.CreateDaduFactory()
	playerFactory := game.CreatePlayerFactory()

	gameFactory := game.CreateGameFactory()
	mygame := gameFactory.CreateGame()
	mygame.SetDaduFactory(daduFactory)
	mygame.SetPlayerFactory(playerFactory)
	mygame.Play(*player, *dadu)
}