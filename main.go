package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var rules = map[string]string{
	"rock":     "scissors",
	"paper":    "rock",
	"scissors": "paper",
}
var gameOptions = []string{"rock", "paper", "scissors"}

type player struct {
	name string
}

type game struct {
	players []player
}

func main() {
	p1 := player{
		name: "Player",
	}

	p2 := player{
		name: "CPU",
	}

	g := game{}
	g.addPlayer(p1)
	g.addPlayer(p2)

	fmt.Printf(`Game starts. Please select one:
1 - %v
2 - %v
3 - %v
0 - Exit program
`, strings.Title(gameOptions[0]),
		strings.Title(gameOptions[1]),
		strings.Title(gameOptions[2]))

	for {
		var option int
		fmt.Print("Enter option: ")
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			fmt.Println("Error while reading input.")
			os.Exit(1)
		}
		if option == 0 {
			fmt.Println("See you next time.")
			os.Exit(1)
		}

		if option > len(gameOptions) {
			fmt.Println("Invalid option. Choose another one")
			continue
		}

		fmt.Println(g.play(gameOptions[option-1]))
	}
}

func rockPaperScissor() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	i := r.Intn(len(gameOptions))
	return gameOptions[i]
}

func (g *game) addPlayer(p player) {
	g.players = append(g.players, p)
}

func (g game) play(playerOption string) string {
	cpuOption := rockPaperScissor()

	var result string
	if playerOption == cpuOption {
		result = "Draw"
	} else if rules[playerOption] == cpuOption {
		result = fmt.Sprintf("%v won", g.players[0].name)
	} else {
		result = fmt.Sprintf("%v won", g.players[1].name)
	}
	return result
}
