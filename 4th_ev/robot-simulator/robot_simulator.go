// Package robot provides 3 step implementation of a robot simulator
package robot

import (
	"errors"
)

// Setp1

func (d Dir) String() string {
	return string(d)
}

const (
	// N - North
	N = Dir(0)
	// E - East
	E = Dir(1)
	// S - South
	S = Dir(2)
	// W - West
	W = Dir(3)
)

// Advance moves forward Step1Robot
func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
		break
	case S:
		Step1Robot.Y--
		break
	case E:
		Step1Robot.X++
		break
	case W:
		Step1Robot.X--
	}
}

// Right turns right Step1Robot
func Right() {
	if Step1Robot.Dir+1 > 3 {
		Step1Robot.Dir = N
	} else {
		Step1Robot.Dir++
	}
}

// Left turns left Step1Robot
func Left() {
	if Step1Robot.Dir-1 < 0 {
		Step1Robot.Dir = W
	} else {
		Step1Robot.Dir--
	}
}

// Step2

// Action models action delegated to a robot
type Action byte

const (
	// R - turn right
	R = Action('R')
	// L - turn left
	L = Action('L')
	// A - advance
	A = Action('A')
	// OFF - turn off
	OFF = Action('O')
)

func isValid(cmd Command) bool {
	return cmd == 'R' || cmd == 'L' || cmd == 'A'
}

// StartRobot translates commands into actions
func StartRobot(cmd <-chan Command, act chan<- Action) {
	for c := range cmd {
		if isValid(c) {
			act <- Action(c)
		}
	}
	// if there is no more commands trigger OFF action
	act <- OFF
}

// Right turns robot right
func (r *Step2Robot) Right() {
	if r.Dir+1 > 3 {
		r.Dir = N
	} else {
		r.Dir++
	}
}

// Left turns robot left
func (r *Step2Robot) Left() {
	if r.Dir-1 < 0 {
		r.Dir = W
	} else {
		r.Dir--
	}
}

// Advance moves robot forward according to current direction
func (r *Step2Robot) Advance() {
	switch r.Dir {
	case N:
		r.Northing++
		break
	case S:
		r.Northing--
		break
	case E:
		r.Easting++
		break
	case W:
		r.Easting--
	}
}

// Room models physics for robot simulator
func Room(rect Rect, robot Step2Robot, act <-chan Action, rep chan<- Step2Robot) {
	for a := range act {
		prevPos := robot.Pos
		switch a {
		case R:
			robot.Right()
			break
		case L:
			robot.Left()
			break
		case A:
			robot.Advance()
			if !isPosInRect(robot.Pos, rect) {
				robot.Pos = prevPos
			}
			break
		case OFF:
			rep <- robot
			break
		}
	}
}

func isPosInRect(pos Pos, rect Rect) bool {
	return pos.Easting >= rect.Min.Easting &&
		pos.Easting <= rect.Max.Easting &&
		pos.Northing >= rect.Min.Northing &&
		pos.Northing <= rect.Max.Northing
}

// Step3

// Action3 models action for particular robot
type Action3 struct {
	RobotName string
	Action
}

// StartRobot3 translates command script for given robot into actions
func StartRobot3(name, script string, action chan Action3, log chan string) {
	for _, c := range script {
		if !isValid(Command(c)) {
			log <- "an undefined command in a script"
			break
		}
		action <- Action3{name, Action(c)}
	}
	// if there is no more commands trigger OFF action
	action <- Action3{name, OFF}
}

// Room3 models physics for robot simulator
func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	if err := validateRobots(robots, extent); err != nil {
		log <- err.Error()
		report <- robots
		return
	}
	byName := indexByName(robots)
	cnt := len(robots)

	for a := range action {
		r, ok := byName[a.RobotName]
		if !ok {
			log <- "an action from an unknown robot"
			report <- robots
			return
		}
		switch a.Action {
		case A:
			if err := tryAdvance(r, robots, extent); err != nil {
				log <- err.Error()
			}
			break
		case R:
			r.Right()
			break
		case L:
			r.Left()
			break
		case OFF:
			if cnt--; cnt == 0 {
				report <- robots
				return
			}
			break
		}
	}
}

func validateRobots(robots []Step3Robot, extent Rect) error {
	usedName := map[string]bool{}
	usedPos := map[Pos]bool{}
	for _, r := range robots {
		if r.Name == "" {
			return errors.New("a robot without a name")
		}
		if usedName[r.Name] {
			return errors.New("duplicate robot names")
		}
		usedName[r.Name] = true
		if usedPos[r.Pos] {
			return errors.New("robots placed at the same place")
		}
		usedPos[r.Pos] = true
		if !isPosInRect(r.Pos, extent) {
			return errors.New("a robot placed outside of the room")
		}
	}
	return nil
}

func indexByName(robots []Step3Robot) map[string]*Step3Robot {
	byName := map[string]*Step3Robot{}
	for idx, r := range robots {
		byName[r.Name] = &robots[idx]
	}
	return byName
}

func tryAdvance(r *Step3Robot, robots []Step3Robot, extent Rect) error {
	prevPos := r.Pos
	r.Advance()
	if !isPosFree(r, robots) {
		r.Pos = prevPos
		return errors.New("a robot attempting to advance into another robot")
	}
	if !isPosInRect(r.Pos, extent) {
		r.Pos = prevPos
		return errors.New("a robot attempting to advance into a wall")
	}
	return nil
}

func isPosFree(r *Step3Robot, robots []Step3Robot) bool {
	for _, other := range robots {
		if r.Name == other.Name {
			continue
		}
		if r.Pos == other.Pos {
			return false
		}
	}
	return true
}
