package main

import (
	"fmt"
	"math/rand"
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

	fmt.Println(g.play())
}

func (player) rockPaperScissor() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	i := r.Intn(len(gameOptions))
	return gameOptions[i]
}

func (g *game) addPlayer(p player) {
	g.players = append(g.players, p)
}

func (g game) play() string {
	p1Chosen := g.players[0].rockPaperScissor()
	p2Chosen := g.players[1].rockPaperScissor()

	var result string
	if p1Chosen == p2Chosen {
		result = "Draw"
	} else if rules[p1Chosen] == p2Chosen {
		result = fmt.Sprintf("%v won", g.players[0].name)
	} else {
		result = fmt.Sprintf("%v won", g.players[1].name)
	}
	return result
}
