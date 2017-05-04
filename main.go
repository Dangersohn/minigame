package main

import "fmt"

func main() {
	//legt units an
	karl := NewUnit(100, 5, 10, true)
	fmt.Println(karl.Angriffe[0])
}
