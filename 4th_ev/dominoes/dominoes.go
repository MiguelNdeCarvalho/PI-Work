package dominoes

type Domino [2]int

func MakeChain(input []Domino) (chain []Domino, ok bool) {
	
	switch len(input) {
	case 0:
		return []Domino{}, true
	case 1:
		

		if input[0][0] == input[0][1] {
			return input, true
		}
		return nil, false
	}
	chain, ok = buildDominoChain(input)
	if !ok {
		return nil, false
	}
	return chain, true
}

func buildDominoChain(d []Domino) ([]Domino, bool) {
	
	permuteList := dominoPermutations(d, len(d))
	for _, p := range permuteList {
		chain, ok := arrangeChain(p)
		if ok {
			return chain, true
		}
	}
	return nil, false
}



func arrangeChain(d []Domino) (chain []Domino, ok bool) {
	

	chain = make([]Domino, len(d))
	n := 0
	

	chain[n] = d[0]
	

	for i := 1; i < len(d); i++ {
		
		t := d[i]
		last := chain[n]
		if n == 0 && last[0] == t[0] {
			

			chain[n] = reverseDomino(last)
			chain[n+1] = t
		} else if n == 0 && last[0] == t[1] {
			

			chain[n] = reverseDomino(last)
			chain[n+1] = reverseDomino(t)
		} else if last[1] == t[0] {
			

			chain[n+1] = t
		} else if last[1] == t[1] {
			

			chain[n+1] = reverseDomino(t)
		} else {
			

			return nil, false
		}
		

		n++
	}
	

	if chain[0][0] != chain[len(chain)-1][1] {
		return nil, false
	}
	return chain, true
}



func reverseDomino(x Domino) Domino {
	
	return Domino{x[1], x[0]}

}



func dominoPermutations(iterable []Domino, r int) (perms [][]Domino) {
	
	pool := iterable
	n := len(pool)

	if r > n {
		return
	}

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	result := make([]Domino, r)
	for i, el := range indices[:r] {
		result[i] = pool[el]
	}

	p := make([]Domino, len(result))
	copy(p, result)
	perms = append(perms, p)

	for n > 0 {
		i := r - 1
		for ; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j++ {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]

				for k := i; k < r; k++ {
					result[k] = pool[indices[k]]
				}

				p = make([]Domino, len(result))
				copy(p, result)
				perms = append(perms, p)

				break
			}
		}

		if i < 0 {
			return
		}

	}
	return
}