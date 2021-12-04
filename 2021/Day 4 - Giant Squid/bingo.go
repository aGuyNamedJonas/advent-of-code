package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	Marked bool
	Value  int
}

type Bingo struct {
	Cells [5][5]*Cell
}

func NewBingoFromString(bingoStr string) *Bingo {
	newBingo := Bingo{}

	rows := strings.Split(bingoStr, "\n")
	nonEmptyRows := []string{}
	for _, row := range rows {
		if len(row) > 0 {
			nonEmptyRows = append(nonEmptyRows, row)
		}
	}

	for rowIndex, row := range nonEmptyRows {
		cellsStr := strings.Fields(row)

		for columnIndex, cell := range cellsStr {
			num, err := strconv.Atoi(cell)
			if err != nil {
				panic("Something went wrong with the number conversion")
			}

			newCell := &Cell{
				Marked: false,
				Value:  num,
			}
			newBingo.Cells[rowIndex][columnIndex] = newCell
		}
	}

	return &newBingo
}

func (bingo *Bingo) MarkNumber(num int) bool {
	foundNumInBingo := false
	for _, row := range bingo.Cells {
		for _, cell := range row {
			if cell.Value == num {
				foundNumInBingo = true
				cell.Marked = true
			}
		}
	}

	return foundNumInBingo
}

func (bingo *Bingo) ContainsBingo() bool {
	for _, row := range bingo.Cells {
		rowBingo := false
		for i, cell := range row {
			if i == 0 {
				rowBingo = cell.Marked
				continue
			}

			rowBingo = rowBingo && cell.Marked
		}

		if rowBingo {
			return true
		}
	}

	for columnIndex, _ := range bingo.Cells[0] {
		columnBingo := false
		for i, _ := range bingo.Cells {
			currentCell := bingo.Cells[i][columnIndex]
			if i == 0 {
				columnBingo = currentCell.Marked
				continue
			}

			columnBingo = columnBingo && currentCell.Marked
		}

		if columnBingo {
			return true
		}
	}

	return false
}

func (bingo *Bingo) MarkNumberAndScore(num int) (bool, int) {
	foundNum := bingo.MarkNumber(num)
	score := 0
	for _, row := range bingo.Cells {
		for _, cell := range row {
			if !cell.Marked {
				score += cell.Value
			}
		}
	}

	score *= num
	return foundNum, score
}

func (bingo *Bingo) ToString() string {
	str := ""
	for _, row := range bingo.Cells {
		for _, cell := range row {
			if cell.Marked {
				str += fmt.Sprintf("(%02d) ", cell.Value)
			} else {
				str += fmt.Sprintf(" %02d  ", cell.Value)
			}
		}
		str += "\n"
	}
	return str
}

type Bingos []*Bingo

func (bingos *Bingos) RemainingBingoCount() int {
	bingoCount := 0
	for _, bingo := range *bingos {
		if bingo.ContainsBingo() {
			bingoCount += 1
		}
	}

	return len(*bingos) - bingoCount
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please specify which gamefile to load:\ne.g.: go run bingo.go game.txt")
		os.Exit(0)
	}

	gameFile := os.Args[1]
	fmt.Printf("Loading game file: %s...\n", gameFile)

	pcMode := false
	if len(os.Args) > 2 && os.Args[2] == "PC_MODE" {
		pcMode = true
	}

	if pcMode {
		fmt.Printf("\n\n~~ EVERYBODY'S A WINNER ~~\n\n")
	} else {
		fmt.Println("Regular mode...")
	}

	fmt.Printf("Loading game file: %s...\n", gameFile)

	gameFileRaw, err := ioutil.ReadFile(gameFile)
	if err != nil {
		fmt.Print(err)
	}
	gameFileStr := string(gameFileRaw)
	contents := strings.Split(gameFileStr, "\n\n")

	moveStrs := contents[0]
	bingoStrs := contents[1:]

	moves := []int{}
	for _, moveStr := range strings.Split(moveStrs, ",") {
		move, err := strconv.Atoi(moveStr)
		if err != nil {
			panic("Something went wrong with the number conversion")
		}
		moves = append(moves, move)
	}

	bingos := Bingos{}
	for _, bingoStr := range bingoStrs {
		newBingo := NewBingoFromString(bingoStr)
		bingos = append(bingos, newBingo)
	}

	fmt.Printf("Loaded %d moves, and %d bingo boards\n", len(moves), len(bingos))

	fmt.Printf("\n\nLet's play bingo!\n\n")
	for _, move := range moves {
		fmt.Printf("Does anyone have a %d?\n", move)
		for _, bingo := range bingos {
			bingoStatusBefore := bingo.ContainsBingo()
			found, score := bingo.MarkNumberAndScore(move)

			bingoStatusAfter := bingo.ContainsBingo()
			bingoStatusChanged := !bingoStatusBefore && bingoStatusAfter

			if pcMode && bingoStatusChanged {
				fmt.Printf("\n\n~BINGO!~ (Score: %d) (%d bingos remaining)\n\n", score, bingos.RemainingBingoCount())
			}

			if (!pcMode && bingo.ContainsBingo()) || (pcMode && bingoStatusChanged && bingos.RemainingBingoCount() == 0) {
				fmt.Printf(`

    \ | | | | | | /
   -- B I N G O !! --
    / | | | | | | \

%s

   ~~ Score: %d ~~`, bingo.ToString(), score)
				fmt.Print("\n")
				os.Exit(0)
			}

			if found {
				fmt.Print("|")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n\n")
	}
}
