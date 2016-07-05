package snakesandladders

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Game     *Game
	Name     *string
	Position int
}

func (g *Game) AddPlayer(name *string) {
	player := Player{
		Name: name,
		Game: g,
	}
	g.Players = append(g.Players, &player)
}

func (p *Player) Move() {
	roll := p.Game.Dice.Roll()
	fmt.Printf("%v rolled a %v and moves from %v to %v\n",
		*p.Name,
		roll,
		p.Position,
		p.Position+roll)
	p.Position += roll

	if v := p.CheckCollision(); v != nil {
		p.Position = v.End
		fmt.Printf("  %v landed on a %v and finished on %v\n",
			*p.Name,
			v.Type(),
			p.Position)
	}
}

func (p *Player) CheckCollision() *Vector {
	for _, ladder := range p.Game.Ladders {
		if p.Position == ladder.Start {
			return ladder
		}
	}

	for _, snake := range p.Game.Snakes {
		if p.Position == snake.Start {
			return snake
		}
	}

	return nil
}

func (p *Player) Win() bool {
	return p.Position >= 100
}

type Vector struct {
	Start int
	End   int
}

func (v *Vector) Direction() int {
	if v.Start < v.End {
		return +1
	}
	return -1
}

func (v *Vector) Type() string {
	if v.Direction() > 0 {
		return "Ladder"
	}

	return "Snake"
}

func (g *Game) SetupBoard() {
	g.Ladders[0] = &Vector{3, 75}
	g.Ladders[1] = &Vector{21, 58}
	g.Ladders[2] = &Vector{58, 83}
	g.Snakes[0] = &Vector{77, 20}
	g.Snakes[1] = &Vector{52, 36}
	g.Snakes[2] = &Vector{33, 8}
}

type Game struct {
	Dice          *Dice
	CurrentPlayer *Player
	Players       []*Player
	Ladders       [3]*Vector
	Snakes        [3]*Vector
}

type Dice struct {
	random *rand.Rand
}

func NewDice() *Dice {
	return &Dice{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (d *Dice) Roll() int {
	return int(d.random.Float64()*12 + 1)
}

func New(p1 *string, p2 *string) *Game {
	game := Game{
		Players: make([]*Player, 0),
		Dice:    NewDice(),
	}

	game.AddPlayer(p1)
	game.AddPlayer(p2)
	game.CurrentPlayer = game.Players[0]
	game.SetupBoard()

	return &game
}

func (g *Game) NextPlayer() {
	if g.CurrentPlayer == g.Players[0] {
		g.CurrentPlayer = g.Players[1]
	} else {
		g.CurrentPlayer = g.Players[0]
	}
}

func (g *Game) Autoplay() {
	for {
		g.CurrentPlayer.Move()

		if g.CurrentPlayer.Win() {
			fmt.Printf("%s is the winner\n", *g.CurrentPlayer.Name)
			return
		}

		g.NextPlayer()
	}
}
