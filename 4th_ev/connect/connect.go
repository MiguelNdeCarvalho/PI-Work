package connect

// ResultOf returns the result of a nice game of HEX.
func ResultOf(field []string) (string, error) {
	for i := 0; i < len(field[0]); i++ {
		if connected(field, i, 0) {
			return string(field[0][i]), nil
		}
	}
	for i := 0; i < len(field); i++ {
		if connected(field, 0, i) {
			return string(field[i][0]), nil
		}
	}
	return "", nil
}

type point struct {
	x, y int
}

func connected(field []string, startX, startY int) bool {
	// Find goal -- if x == 0 && y > 0 -> x == len(field[0]) is the goal -> the other side;
	// if x > 0 && y == 0 -> y == len(field) -> the bottom is the goal
	seen := make(map[point]struct{})
	start := point{x: startX, y: startY}
	seen[start] = struct{}{}
	path := []point{start}
	var c point
	player := field[startY][startX]
	if player == '.' {
		return false
	}
	for len(path) > 0 {
		c, path = path[0], path[1:]
		// check if we arrived at the goal
		if startX == 0 && startY >= 0 {
			// we are going down the front column, check if we arrived at the other side
			if c.x == len(field[0])-1 {
				return rune(field[c.y][c.x]) == rune(player)
			}
		} else if startX >= 0 && startY == 0 {
			// we are going down the top row, check if arrived at the bottom
			if c.y == len(field)-1 {
				return rune(field[c.y][c.x]) == rune(player)
			}
		}

		moves := move(field, c, rune(player))
		for _, m := range moves {
			if _, ok := seen[m]; !ok {
				path = append(path, m)
				seen[m] = struct{}{}
			}
		}
	}
	return false
}

func move(field []string, c point, player rune) (points []point) {
	if c.x-1 >= 0 && rune(field[c.y][c.x-1]) == player {
		points = append(points, point{x: c.x - 1, y: c.y})
	}
	if c.x+1 < len(field[c.y]) && rune(field[c.y][c.x+1]) == player {
		points = append(points, point{x: c.x + 1, y: c.y})
	}
	if c.x-1 >= 0 && c.y+1 < len(field) && rune(field[c.y+1][c.x-1]) == player {
		points = append(points, point{x: c.x - 1, y: c.y + 1})
	}
	if c.y+1 < len(field) && rune(field[c.y+1][c.x]) == player {
		points = append(points, point{x: c.x, y: c.y + 1})
	}
	if c.y-1 >= 0 && c.x+1 < len(field[c.y-1]) && rune(field[c.y-1][c.x+1]) == player {
		points = append(points, point{x: c.x + 1, y: c.y - 1})
	}
	if c.y-1 >= 0 && rune(field[c.y-1][c.x]) == player {
		points = append(points, point{x: c.x, y: c.y - 1})
	}
	return
}