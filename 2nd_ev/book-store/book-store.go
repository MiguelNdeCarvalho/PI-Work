package main

import(
	"fmt"
	"math"
 	"sort"
 )

type kv struct {
	Key   int
	Value int
}

func Cost(books []int) int {
	cost := math.MaxInt32
	for i := 5; i > 2; i--{
		rows, maxRowLen := getAllRows(books, i)
		rows_cost := discountRules(rows)
		if rows_cost < cost{
			cost = rows_cost
		}
		if maxRowLen != i{ // skip unnecessary loops
			i = maxRowLen
		}
	}
	return cost
}

func createTree(books []int) *map[int]int {
	tree := map[int]int{}
	for _, val := range books {
		tree[val]++
	}
	return &tree
}

func getNextRow(sort_books *[]kv, max_len int) []int {
	res := make([]int, 0)
	for i, info := range *sort_books {
		if info.Value > 0 && len(res) < max_len {
			res = append(res, info.Key)
			(*sort_books)[i].Value--
		}
	}
	return res
}

func sortBooks(tree *map[int]int) *[]kv {
	var sorted_books []kv
	for k, v := range *tree {
		sorted_books = append(sorted_books, kv{k, v})
	}

	sort.Slice(sorted_books, func(i, j int) bool {
		return sorted_books[i].Value > sorted_books[j].Value
	})
	return &sorted_books
}

func getAllRows(books []int, max_len int) ([][]int, int) {
	//worse case - all book uniq
	res := make([][]int, len(books)) 
	idx := 0
	max_row_len := 0
	tree := createTree(books)
	sorted_books := sortBooks(tree)
	for {
		row := getNextRow(sorted_books, max_len)
		if len(row) > 0 {
			res[idx] = row
			idx++
			if max_row_len < len(row) {
				max_row_len = len(row)
			}
		} else {
			return res, max_row_len
		}

	}
}

func discountRules(books [][]int) int {
	var sum int = 0
	for _, rows := range books {
		switch len(rows) {
		case 1:
			sum += 800
		case 2:
			sum += 2 * 800 * 0.95
		case 3:
			sum += 3 * 800 * 0.90
		case 4:
			sum += 4 * 800 * 0.80
		case 5:
			sum += 5 * 800 * 0.75
		}
	}
	return sum
}

var teste = []int{1,1,2,2,3,3,4,5}
func main(){

	fmt.Printf("%d\n",Cost(teste))
}