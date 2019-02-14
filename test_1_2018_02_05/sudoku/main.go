package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const sudokuUsage = `
Hello, it's my programm for sudoku solvening.
DITS test work. Backtracking algorithm.

This is flags for usage:
	-s string - sudoku string for solve.
	-i string - input filename with sudoku string
	-o string - output filename for sudoku solved table
	-d boolean - type demo for solve sudokuDemoTable from .go file
	-f string - SOLVE ALL SUDOKU FROM FOLDER
Also you can combine flags like -i question.txt -o answer.txt - get
from question.txt sudoku string and save solved sudoku in answer.txt
or -s *some_sudoku_string* -d will demonstrate solve with this string.

Note: All sudoku strings which you select with flags added to stack
and will be run at combination. You can run with -s flag and
add -f folder, and they will added to one stack and programm will solve it.

Note 2: Type of sudoku input is sudoku[col][row].
Example: 800000000007500009030000180060001050009040000000750000002070004000003610000000800

 - - - - - - - - - - - -
| 8 . . | . . . | . . . |
| . . 3 | 6 . . | . . . |
| . 7 . | . 9 . | 2 . . |
 - - - - - - - - - - - -
| . 5 . | . . 7 | . . . |
| . . . | . 4 5 | 7 . . |
| . . . | 1 . . | . 3 . |
 - - - - - - - - - - - -
| . . 1 | . . . | . 6 8 |
| . . 8 | 5 . . | . 1 . |
| . 9 . | . . . | 4 . . |
 - - - - - - - - - - - -

`

var (
	sudokuDemoTable = [][]uint8{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}
	sudokuTable   []string
	writter       = os.Stdout
	workTime      time.Duration
	sudokuCounter int
)

func main() {
	sudokuQuestion := flag.String("s", "", "Usage: sudoku.exe -s 009805060706240300020000074600504013050010080190306002260000030007053609010702500")
	inputFilename := flag.String("i", "", `Usage: sudoku.exe -i "C:\test.txt" or test.txt if in same folder with programm`)
	outputFilename := flag.String("o", "", `Usage: sudoku.exe -o "C:\test.txt" or test.txt if in same folder with programm`)
	fromDirPath := flag.String("f", "", `Usage: sudoku.exe -f "C:\sudoku\maps" or .\maps or if you have a spaces in path use "C:\sudoku for real\maps"`)
	isDemo := flag.Bool("d", false, "Usage: sudoku.exe -d. Demonstration work with in-build sudoku in .go file")
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Fprint(os.Stderr, sudokuUsage)
		return
	}
	if *isDemo {
		s, err := createStringFromTable(&sudokuDemoTable)
		if err != nil {
			log.Fatal(err)
		}
		sudokuTable = append(sudokuTable, s)
	}
	if len(*fromDirPath) > 0 {
		if err := getSudokuStringsFromDirectory(*fromDirPath); err != nil {
			log.Fatal(err)
		}
	}
	if len(*inputFilename) > 0 {
		s, err := ioutil.ReadFile(*inputFilename)
		if err != nil {
			log.Fatal(err)
		}
		sl := strings.Split(string(s), "\r\n")
		sudokuTable = append(sudokuTable, sl...)
	}
	if len(*sudokuQuestion) > 0 {
		sudokuTable = append(sudokuTable, *sudokuQuestion)
	}
	if len(*outputFilename) > 0 {
		fileWritter, err := os.Create(*outputFilename)
		if err != nil {
			log.Fatal(err)
		}
		writter = fileWritter
		defer fileWritter.Close()
	}
	for _, table := range sudokuTable {
		if table == "" {
			continue
		}
		if err := startSolveSudoku(table); err != nil {
			log.Fatal(err)
		}
		sudokuCounter++
	}
	exitMessage()
}

func getSudokuStringsFromDirectory(dir string) error {
	d, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range d {
		if !file.IsDir() {
			s, err := ioutil.ReadFile(dir + `\` + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			sl := strings.Split(string(s), "\r\n")
			sudokuTable = append(sudokuTable, sl...)
		}
	}
	return nil
}

func exitMessage() {
	fmt.Fprintf(writter, `
	Solved for %s.
	%d Sudoku.

────────────¸,o​¤°“°¤o,¸────────────
─────────── (….◕​ ‿ ◕.…) ───────────
───────── oOO——"​♥""——OOo ───────────
▀█▀─█▄█─█▀█─█▄─█─█▄▀──█▄█─█▀█─█─█
─█──█▀█─█▀█─█─▀█─█▀▄───█──█▄█─█▄█
`, workTime, sudokuCounter)
}
