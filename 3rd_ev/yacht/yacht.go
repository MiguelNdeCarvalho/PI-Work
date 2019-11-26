package yacht

func Score(dice []int, category string) int {
	var counter [6]int

	for _, n := range dice {
		counter[n-1]++
	}

	switch category {
	case "ones":
		return counter[0]
	case "twos":
		return counter[1] * 2
	case "threes":
		return counter[2] * 3
	case "fours":
		return counter[3] * 4
	case "fives":
		return counter[4] * 5
	case "sixes":
		return counter[5] * 6
	case "full house":
		var (
			score     int
			fullHouse int
		)
		for i := range counter {
			switch counter[i] {
			case 2, 3:
				fullHouse += counter[i]
				score += counter[i] * (i + 1)
			}
		}
		if fullHouse == 5 {
			return score
		}
	case "four of a kind":
		for i := range counter {
			if counter[i] > 3 {
				return 4 * (i + 1)
			}
		}
	case "little straight":
		for i := 0; i < 5; i++ {
			if counter[i] != 1 {
				return 0
			}
		}
		return 30
	case "big straight":
		for i := 1; i < 6; i++ {
			if counter[i] != 1 {
				return 0
			}
		}
		return 30
	case "choice":
		var sum int
		for i := range dice {
			sum += dice[i]
		}
		return sum
	case "yacht":
		for i := range counter {
			if counter[i] == 5 {
				return 50
			}
		}
	}

	return 0
}
