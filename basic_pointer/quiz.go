package main

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Zelda"}

	gamer.AddGame("Mario")
	gamer.AddGame("Bioshock")

	for _, game := range gamer.Games {
		println("- ", game)
	}
}