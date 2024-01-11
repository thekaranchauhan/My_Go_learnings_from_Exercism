package bowling

import "errors"

type Game struct {
	throws   []int
	isSecond bool
}

func NewGame() *Game {
	return &Game{throws: make([]int, 0, 21)}
}

func (g *Game) Roll(pins int) error {
	pinsDown := 0
	if g.isSecond {
		pinsDown = g.throws[len(g.throws)-1]
	}
	if _, err := g.Score(); err == nil || pins < 0 || pins+pinsDown > 10 {
		return errors.New("invalid roll")
	}
	g.throws, g.isSecond = append(g.throws, pins), !g.isSecond
	if pins == 10 {
		g.isSecond = false
	}
	return nil
}

func (g *Game) Score() (total int, err error) {
	for throw, frame := 0, 0; frame < 10; throw, frame = throw+1, frame+1 {
		if len(g.throws) <= throw+1 {
			return 0, errors.New("not enough throws")
		}
		one, two := g.throws[throw], g.throws[throw+1]
		total += one + two
		if one == 10 || one+two == 10 {
			if len(g.throws) <= throw+2 {
				return 0, errors.New("not enough throws")
			}
			total += g.throws[throw+2]
		}
		if one != 10 {
			throw++
		}
	}
	return total, nil
}
