package gameoflife

import ()

type Game struct {
	current int
	grids   [][][]int
}

func NewGame(seed [][]int) *Game {
	game := new(Game)
	game.current = 0
	game.grids = make([][][]int, 2, 2)
	for x := 0; x < 2; x++ {
		game.grids[x] = make([][]int, len(seed), len(seed))
		for i := 0; i < len(seed); i++ {
			game.grids[x][i] = make([]int, len(seed[i]), len(seed[i]))
		}
	}
	for i := 0; i < len(seed); i++ {
		for y := 0; y < len(seed[i]); y++ {
			game.grids[0][i][y] = seed[i][y]
		}
	}
	return game
}

func (game *Game) Grid() [][]int {
	return game.grids[game.current]
}

func (game *Game) NextGen() {
	var next int
	if game.current == 0 {
		next = 1
	} else {
		next = 0
	}
	for y := 0; y < len(game.grids[next]); y++ {
		for x := 0; x < len(game.grids[next][y]); x++ {
			game.grids[next][y][x] = getCellValue(y, x, game.grids[game.current])
		}
	}
	game.current = next
}

func getCellValue(y, x int, grid [][]int) int {
	adj := getAdjacentCells(y, x, grid)
	switch {
	case adj < 2:
		return 0
	case adj == 3:
		return 1
	case adj > 3:
		return 0
	default:
		return grid[y][x]
	}
}

func getAdjacentCells(y, x int, grid [][]int) int {
	lastY := y - 1
	nextY := y + 1
	lastX := x - 1
	nextX := x + 1
	if lastY < 0 {
		lastY = len(grid) - 1
	}
	if nextY > len(grid)-1 {
		nextY = 0
	}
	if lastX < 0 {
		lastX = len(grid[y]) - 1
	}
	if nextX > len(grid[y])-1 {
		nextX = 0
	}
	total := 0
	total += grid[lastY][lastX] + grid[lastY][x] + grid[lastY][nextX]
	total += grid[y][lastX] + grid[y][nextX]
	total += grid[nextY][lastX] + grid[nextY][x] + grid[nextY][nextX]
	return total
}
