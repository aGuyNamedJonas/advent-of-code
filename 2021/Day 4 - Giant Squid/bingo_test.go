package main

import "testing"

func TestNewBingoFromInput(t *testing.T) {
	bingoStr := `
		84 94 24 52 44
		96 33 74 35 13
		60 51 41 19 95
		50 93 27 40  1
		67 23 37 88 85
	`

	bingo := NewBingoFromString(bingoStr)
	middleCell := bingo.Cells[2][2].Value
	if middleCell != 41 {
		t.Errorf("Middle cell in bingo expected to be 41, instead: %d", middleCell)
	}

	middleCellMarked := bingo.Cells[2][2].Marked
	if middleCellMarked != false {
		t.Errorf("Middle cell in bingo expected to be unmarked, instead: %t", middleCellMarked)
	}
}

func TestMarkNumber(t *testing.T) {
	bingoStr := `
		84 94 24 52 44
		96 33 74 35 13
		60 51 41 19 95
		50 93 27 40  1
		67 23 37 88 85
	`

	bingo := NewBingoFromString(bingoStr)
	found123 := bingo.MarkNumber(123)
	if found123 {
		t.Error("Found 123, even though it's not part of the bingo")
	}

	found41 := bingo.MarkNumber(41)
	if !found41 {
		t.Error("Did not find 41, even though it's in the bingo")
	}

	middleCellMarked := bingo.Cells[2][2].Marked
	if !middleCellMarked {
		t.Error("Middle cell not marked, even though it should be (it's the number 41)")
	}
}

func TestContainsBingo(t *testing.T) {
	bingoStr := `
		84 94 24 52 44
		96 33 74 35 13
		60 51 41 19 95
		50 93 27 40  1
		67 23 37 88 85
	`

	containsRowBingo := NewBingoFromString(bingoStr)
	containsRowBingo.MarkNumber(60)
	containsRowBingo.MarkNumber(51)
	containsRowBingo.MarkNumber(41)
	containsRowBingo.MarkNumber(19)
	if containsRowBingo.ContainsBingo() {
		t.Error("Row Bingo should not yet contain a bingo")
	}

	containsRowBingo.MarkNumber(95)
	if !containsRowBingo.ContainsBingo() {
		t.Error("Row Bingo should have a bingo")
	}

	containsColumnBingo := NewBingoFromString(bingoStr)
	containsColumnBingo.MarkNumber(24)
	containsColumnBingo.MarkNumber(74)
	containsColumnBingo.MarkNumber(41)
	containsColumnBingo.MarkNumber(27)
	if containsColumnBingo.ContainsBingo() {
		t.Error("Column Bingo should not yet contain a bingo")
	}

	containsColumnBingo.MarkNumber(37)
	if !containsColumnBingo.ContainsBingo() {
		t.Error("Column Bingo should have a bingo")
	}
}

func TestMarkNumberAndScore(t *testing.T) {
	bingoStr := `
		14 21 17 24  4
		10 16 15  9 19
		18  8 23 26 20
		22 11 13  6  5
		2  0 12  3  7
	`

	bingo := NewBingoFromString(bingoStr)
	bingo.MarkNumber(7)
	bingo.MarkNumber(4)
	bingo.MarkNumber(9)
	bingo.MarkNumber(5)
	bingo.MarkNumber(11)
	bingo.MarkNumber(17)
	bingo.MarkNumber(23)
	bingo.MarkNumber(2)
	bingo.MarkNumber(0)
	bingo.MarkNumber(14)
	bingo.MarkNumber(21)

	found, score := bingo.MarkNumberAndScore(24)
	if !found {
		t.Error("Expected 4 to be found")
	}

	if score != 4512 {
		t.Errorf("Expected score to be 4512, instead go %d\n", score)
	}
}

func TestRemainingBingoCount(t *testing.T) {
	bingoStr := `
	84 94 24 52 44
	96 33 74 35 13
	60 51 41 19 95
	50 93 27 40  1
	67 23 37 88 85
	`

	bingo := NewBingoFromString(bingoStr)
	bingos := Bingos{bingo}

	bingo.MarkNumber(60)
	bingo.MarkNumber(51)
	bingo.MarkNumber(41)
	bingo.MarkNumber(19)
	if bingos.RemainingBingoCount() != 1 {
		t.Errorf("Remainig bingo count should be 1, instead go %d", bingos.RemainingBingoCount())
	}

	bingo.MarkNumber(95)
	if bingos.RemainingBingoCount() != 0 {
		t.Errorf("Remainig bingo count should be 0, instead go %d", bingos.RemainingBingoCount())
	}
}
