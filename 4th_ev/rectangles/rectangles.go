package rectangles

func Search(x, y, largura, altura int, diagrama []string) int {
	count := 0
	for xx := x + 1; xx < largura && (diagrama[y][xx] == '-' || diagrama[y][xx] == '+'); xx++ {
		
		if diagrama[y][xx] != '+' {
			continue
		}

		for yy := y + 1; yy < altura && (diagrama[yy][xx] == '|' || diagrama[yy][xx] == '+') && (diagrama[yy][x] == '|' || diagrama[yy][x] == '+'); yy++ {
			if diagrama[yy][xx] == '+' && diagrama[yy][x] == '+' {
				
				xxx := x + 1

				for ; xxx < xx && (diagrama[yy][xxx] == '-' || diagrama[yy][xxx] == '+'); xxx++ {
				}
				if xxx == xx {
					count++
				}
			}
		}
	}
	return count
}

// Conta os retangulos na tabela ASCII.
func Count(diagrama []string) int {

	
	altura := len(diagrama)

	if altura == 0 {
		return 0
	}

	largura, count := len(diagrama[0]), 0

	
	for y := 0; y < altura-1; y++ {
		for x := 0; x < largura-1; x++ {
			if diagrama[y][x] == '+' {
				
				count += Search(x, y, largura, altura, diagrama)
			}
		}
	}
	return count
}