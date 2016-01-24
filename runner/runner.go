package main

import (
	"bufio"
	"fmt"
	"github.com/davidafox/gameoflife"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	file, err := os.Open("testfile.life")
	if err != nil {
		log.Println(err)
		return
	}
	p := NewParser("0", ".")
	grid := p.Parse(file)
	game := gameoflife.NewGame(grid)
	for true {
		display(game.Grid())
		game.NextGen()
		clear := exec.Command("cls")
		clear.Stdout = os.Stdout
		//		err = clear.Run()
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second / 10)
		fmt.Println("\n\n\n")
	}
}

type Parser struct {
	liveCharacter string
	deadCharacter string
}

func (p *Parser) lineParser(line string) []int {
	result := make([]int, len(line))
	for i := 0; i < len(line); i++ {
		switch string(line[i]) {
		case p.liveCharacter:
			result[i] = 1
		case p.deadCharacter:
			result[i] = 0
		default:
			log.Println("invalidCharacter: ", line[i])
		}
	}
	return result
}

func (p *Parser) Parse(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)
	result := make([][]int, 0)
	for i := 0; scanner.Scan(); i++ {
		result = append(result, p.lineParser(scanner.Text()))
	}
	return result
}

func NewParser(liveCharacter, deadCharacter string) *Parser {
	p := new(Parser)
	p.liveCharacter = liveCharacter
	p.deadCharacter = deadCharacter
	return p
}

func display(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		result := ""
		for y := 0; y < len(grid[i]); y++ {
			if grid[i][y] == 1 {
				result += "0"
			} else {
				result += "."
			}
		}
		fmt.Println(result)
	}
}
