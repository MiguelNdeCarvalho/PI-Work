package bowling

import (
	"errors"
	"strings"
)

const testVersion = 1

type Frame struct {
	pins []int
	roll int
}

func (f *Frame) IsStrike() bool {
	return f.pins[0] == 10
}

func (f *Frame) IsSpare() bool {
	return f.pins[0] < 10 && f.pins[0]+f.pins[1] == 10
}

func (f *Frame) Extend() {
	tmp := make([]int, 3)
	tmp[0] = f.pins[0]
	tmp[1] = f.pins[1]
	f.pins = tmp
}

func (f *Frame) Score() int {
	score := 0
	for _, p := range f.pins {
		score += p
	}
	return score
}

func (f *Frame) String() string {
	rolls := make([]string, len(f.pins))
	for i, p := range f.pins {
		if p == 10 {
			rolls[i] = "X"
		} else if i == 1 && f.IsSpare() {
			rolls[i] = "/"
		} else {
			rolls[i] = string('0' + p)
		}
	}
	return strings.Join(rolls, " ")
}

type Game struct {
	frame  int
	frames []Frame
}

func NewGame() *Game {
	frames := make([]Frame, 11)
	for i := 1; i < 11; i++ {
		frames[i] = Frame{make([]int, 2), 0}
	}
	return &Game{1, frames}
}

func (g *Game) String() string {
	fs := make([]string, len(g.frames)-1)
	for i := 0; i < 10; i++ {
		fs[i] = g.frames[i+1].String()
	}
	return strings.Join(fs, " | ")
}

func (g *Game) Roll(pins int) error {
	if g.frame == 11 {
		return errors.New("cannot roll when game is over")
	}
	if pins < 0 || pins > 10 {
		return errors.New("invalid pin count")
	}
	frame := &g.frames[g.frame]
	frame.pins[frame.roll] = pins
	if (frame.pins[0]+frame.pins[1] > 10 &&
		(g.frame != 10 ||
			!frame.IsStrike())) ||
		(len(frame.pins) == 3 &&
			!frame.IsSpare() &&
			frame.pins[1] < 10 &&
			frame.pins[1]+frame.pins[2] > 10) {
		return errors.New("invalid pin count")
	}
	if g.frame == 10 {
		if (frame.roll == 0 && frame.IsStrike()) ||
			(frame.roll == 1 && frame.IsSpare()) {
			frame.Extend()
			frame.roll++
		} else if frame.roll == len(frame.pins)-1 {
			g.frame++
		} else {
			frame.roll++
		}
	} else if frame.IsStrike() || frame.roll == len(frame.pins)-1 {
		g.frame++
	} else {
		frame.roll++
	}
	return nil
}

func (g *Game) Score() (int, error) {
	if g.frame != 11 {
		return 0, errors.New("cannot score until game end")
	}
	score := 0
	for i := 1; i < 11; i++ {
		f := g.frames[i]
		score += f.Score()
		if i < 10 {
			if f.IsStrike() {
				score += g.frames[i+1].Score()
				if i+2 < 10 && g.frames[i+1].IsStrike() {
					score += g.frames[i+2].pins[0]
				}
			} else if f.IsSpare() {
				score += g.frames[i+1].pins[0]
			}
		}
	}
	return score, nil
}
