package main

import "fmt"

// to run use go run *.go
func main() {
	sl := CreateCacheSlice()
	sl.AddToSlice("1", "2", "3", "4", "5", "6")
	fmt.Println(sl.Delete(3))
	fmt.Println(sl)
}
