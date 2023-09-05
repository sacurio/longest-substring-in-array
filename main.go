package main

import (
	"fmt"
)

const (
	Horizontal = "--"
	Vertical   = "|"
	LDiagonal  = `/`
	RDiagonal  = `\`
)

type (
	subsequent struct {
		character   string
		direction   string
		repetitions int
		dataset     map[string]int
	}
)

func NewSubsequent(character, direction string, repetitions int) *subsequent {
	return &subsequent{
		character:   character,
		direction:   direction,
		repetitions: repetitions,
		dataset:     make(map[string]int),
	}
}

func main() {
	initialize()
}

func findConsecutiveStrings(dir string, path []string) (map[string]int, string, int) {
	var max int
	var m string
	collector := make(map[string]int)
	for i := 1; i < len(path); i++ {
		if path[i] == path[i-1] {
			if _, ok := collector[path[i]]; !ok {
				collector[path[i]] += 2
			} else {
				collector[path[i]] += 1
			}
			if collector[path[i]] > max {
				max = collector[path[i]]
				m = path[i]
			}
		}
	}
	return collector, m, max
}

func (s *subsequent) maxValue(max *int, sb *subsequent, data []string) {
	ds, character, repetitions := findConsecutiveStrings(Horizontal, data)

	if repetitions > *max {
		sb.character = character
		sb.direction = s.direction
		sb.repetitions = repetitions
		sb.dataset = ds
		*max = repetitions
	}
}

func breakInPaths(data [][]string) {
	horizontal := make(map[int][]string)
	vertical := make(map[int][]string)
	leftDiagonal := make(map[int][]string)
	rightDiagonal := make(map[int][]string)
	maxSubsequent := &subsequent{}
	var max int
	for i, row := range data {
		horizontal[i] = row
		sb := &subsequent{direction: Horizontal}
		sb.maxValue(&max, maxSubsequent, row)
		for j, column := range row {
			vertical[j] = append(vertical[j], column)
			rightDiagonal[i+1+j] = append(rightDiagonal[i+1+j], row[len(row)-1-j])
			leftDiagonal[i+1+j] = append(leftDiagonal[i+1+j], column)
		}
	}

	for i := 0; i < len(rightDiagonal); i++ {
		if i < len(vertical) {
			sb := &subsequent{direction: Vertical}
			sb.maxValue(&max, maxSubsequent, vertical[i])
		}
		sb := &subsequent{direction: LDiagonal}
		sb.maxValue(&max, maxSubsequent, leftDiagonal[i])

		sb = &subsequent{direction: RDiagonal}
		sb.maxValue(&max, maxSubsequent, rightDiagonal[i])
	}

	fmt.Printf("Los caracteres consecuentes más repetidos son: %s, en un número de %d repeticiones y en sentido %s\n", maxSubsequent.character, maxSubsequent.repetitions, maxSubsequent.direction)
}

func initialize() {
	var n int
	fmt.Print("Ingrese el tamaño de la matriz cuadrada: ")
	fmt.Scan(&n)

	matriz := make([][]string, n)
	for i := range matriz {
		matriz[i] = make([]string, n)
	}

	fmt.Println("Ingrese los elementos de la matriz:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("Elemento [%d][%d]: ", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}

	fmt.Println("Matriz ingresada:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%s\t", matriz[i][j])
		}
		fmt.Println()
	}
	breakInPaths(matriz)
}
