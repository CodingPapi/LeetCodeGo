package pureGo

import "math"

var store [][][]int
var MOD = int(math.Pow10(9)) + 7

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	store = make([][][]int, m)
	for i := range store {
		store[i] = make([][]int, n)
		for j := range store[i] {
			store[i][j] = make([]int, maxMove+1)
			for k := 0; k <= maxMove; k++ {
				store[i][j][k] = -1
			}
		}
	}
	result := doRe(m, n, maxMove, startRow, startColumn)
	return result
}

func doRe(m int, n int, maxMove int, startRow int, startColumn int) int {
	if startRow < 0 || startColumn < 0 || startRow >= m || startColumn >= n {
		return 1
	}
	if maxMove <= 0 {
		return 0
	}
	if store[startRow][startColumn][maxMove] != -1 {
		return store[startRow][startColumn][maxMove]
	}
	if startRow-maxMove >= 0 && startRow+maxMove < m && startColumn-maxMove >= 0 && startColumn+maxMove < n {
		store[startRow][startColumn][maxMove] = 0
		return 0
	}
	maxMove--
	currPaths := doRe(m, n, maxMove, startRow-1, startColumn) % MOD
	currPaths = (currPaths + doRe(m, n, maxMove, startRow+1, startColumn)) % MOD
	currPaths = (currPaths + doRe(m, n, maxMove, startRow, startColumn-1)) % MOD
	currPaths = (currPaths + doRe(m, n, maxMove, startRow, startColumn+1)) % MOD
	store[startRow][startColumn][maxMove+1] = currPaths
	return currPaths
}
