package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var countIteration int
var printVals = map[uint8]string{
	0: ". ",
	1: "1 ",
	2: "2 ",
	3: "3 ",
	4: "4 ",
	5: "5 ",
	6: "6 ",
	7: "7 ",
	8: "8 ",
	9: "9 ",
}

func startSolveSudoku(sudoku string) error {
	countIteration = 0
	table, err := createTableFromString(sudoku)
	if err != nil {
		return err
	}
	printTable("\nYour sudoku table", table)
	start := time.Now()
	if !solveSudoku(table) {
		return fmt.Errorf("\nYour sudoku table incorrent")
	}
	sudokuTime := time.Now().Sub(start)
	workTime += sudokuTime
	printTable("\nSolved sudoku table", table)
	fmt.Fprintf(writter, "Solved for %s.\nWith %d iterations.\n", sudokuTime, countIteration)
	fmt.Fprintln(writter, "*************************")
	return nil
}

func createStringFromTable(array *[][]uint8) (string, error) {
	s := bytes.NewBufferString("")
	for rowIndex := range *array {
		for colIndex := range *array {
			if _, err := s.WriteString(strconv.Itoa(int((*array)[colIndex][rowIndex]))); err != nil {
				return "", nil
			}
		}
	}
	return s.String(), nil
}

func createTableFromString(sudoku string) (*[][]uint8, error) {
	if sudoku == "" {
		return nil, fmt.Errorf("Invalid Sudoku string, wrong size %d", len(sudoku))
	}
	size := int(math.Sqrt(float64(len(sudoku))))
	if len(sudoku)%size != 0 {
		return nil, fmt.Errorf("Invalid Sudoku string, wrong size %d", len(sudoku))
	}
	array := make([][]uint8, size)
	for row := range array {
		array[row] = make([]uint8, size)
	}
	colIndex, rowIndex := 0, 0
	for idx, num := range sudoku {
		val, err := strconv.Atoi(string(num))
		if err != nil {
			return nil, fmt.Errorf("Invalid Sudoku string, can't parse %q", num)
		}
		if (idx+1)%size == 0 && idx != 0 {
			array[rowIndex][colIndex] = uint8(val)
			colIndex++
			rowIndex = 0
		} else {
			array[rowIndex][colIndex] = uint8(val)
			rowIndex++
		}
	}
	return &array, nil
}

func isCorrect(row, col int, num uint8, array *[][]uint8) bool {
	for idx := range *array {
		if (*array)[row][idx] == num {
			return false
		}
		if (*array)[idx][col] == num {
			return false
		}
	}

	sqrt := int(math.Sqrt(float64(len(*array))))
	boxRowStart := row - row%sqrt
	boxColStart := col - col%sqrt

	for r := boxRowStart; r < boxRowStart+sqrt; r++ {
		for d := boxColStart; d < boxColStart+sqrt; d++ {
			if (*array)[r][d] == num {
				return false
			}
		}
	}
	return true
}

func solveSudoku(array *[][]uint8) bool {
	countIteration++
	length := len(*array)
	row := -1
	col := -1
	isEmpty := true
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if (*array)[i][j] == 0 {
				row = i
				col = j
				isEmpty = false
				break
			}
		}
		if !isEmpty {
			break
		}
	}
	if isEmpty {
		return true
	}
	for num := 1; num <= length; num++ {
		if isCorrect(row, col, uint8(num), array) {
			(*array)[row][col] = uint8(num)
			if solveSudoku(array) {
				return true
			}
			(*array)[row][col] = 0
		}
	}
	return false
}

func printTable(label string, array *[][]uint8) {
	const line = " - - - -"
	size := int(math.Sqrt(float64(len(*array))))
	fmt.Fprintln(writter, label)
	for i := range *array {
		if i%size == 0 {
			fmt.Fprintln(writter, strings.Repeat(line, size))
		}
		for j := range *array {
			if j%size == 0 {
				fmt.Fprint(writter, "| ")
			}
			fmt.Fprint(writter, printVals[(*array)[i][j]])
		}
		fmt.Fprintln(writter, "|")
	}
	fmt.Fprintln(writter, strings.Repeat(line, size))
}
