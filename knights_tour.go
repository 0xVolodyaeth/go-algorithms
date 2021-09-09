package main

import "fmt"

const (
	NumRows = 8
	NumCols = 8
)

func isSave(r, c int, moveNumber [][]int) bool {
	if (r >= 0) && (r < NumRows) && (c >= 0) && (c < NumCols) && moveNumber[r][c] == 0 {
		return true
	}

	return false
}

func printSolution(solution [][]int) {
	for i := 0; i < len(solution); i++ {
		for k := 0; k < len(solution); k++ {
			fmt.Printf("%d ", solution[i][k])
		}
		fmt.Println()
	}
}

// Move the knight to position [row, col]. Then recursively try
// to make other moves. Return true if we find a valid solution.
func KnightsTour(row int, col int, moveNumber [][]int, numMovestaken int) bool {

	// Move the knight to this position
	numMovestaken++
	moveNumber[row][col] = numMovestaken

	fmt.Println(numMovestaken)

	// See if we have made all required moves
	if numMovestaken == 64 {
		return true
	}

	// Build arrays to determine where legal moves are
	// with respect to this position.
	dRows := []int{2, 1, -1, -2, -2, -1, 1, 2}
	dCols := []int{1, 2, 2, 1, -1, -2, -2, -1}

	// for i := 0; i < 8; i++ {
	// 	r := row + dRows[i]
	// 	c := row + dCols[i]

	// 	if isSave(r, c, moveNumber) {

	// 		if KnightsTour(r, c, moveNumber, numMovestaken) {
	// 			return true
	// 		}

	// 		moveNumber[r][c] = 0
	// 	}
	// }

	solution := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	if solveUtil(row, col, 1, solution, dRows, dCols) {
		print("solution is found\n")
		printSolution(solution)
		return true
	}
	print("solution is not found")

	return false
}

func solveUtil(x, y int, moveNum int, solution [][]int, xMove, yMove []int) bool {

	if moveNum == 64 {
		return true
	}

	for k := 0; k < 8; k++ {
		nextX := x + xMove[k]
		nextY := y + yMove[k]

		if isSave(nextX, nextY, solution) {
			solution[nextX][nextY] = moveNum
			if solveUtil(nextX, nextY, moveNum+1, solution, xMove, yMove) {
				return true
			}
			solution[nextX][nextY] = 0
		}
	}

	return false

}

func main() {

	board := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	// board6 := [][]int{
	// 	{0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0},
	// 	{0, 0, 0, 0, 0, 0},
	// }

	fmt.Println(KnightsTour(0, 0, board, 0))

}
