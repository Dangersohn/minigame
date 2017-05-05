package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Menu gibt den startus der Einheiten aus und die Atacken
func Menu(you *Unit, enemy *Unit) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Life: %d \t Enemy Life: %d\n", you.Leben, enemy.Leben)
	Atkausgabe(you)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	switch input {
	case "1":
		SchlagNormal(you, enemy)
	case "2":
		SchlagSchwer(you, enemy)
	}
}

//Atkausgabe test blub
func Atkausgabe(you *Unit) {
	for i := 0; i < len(you.Angriffe); i++ {
		if you.Angriffe[i] != "" {
			fmt.Printf("%d) %s \n", i+1, you.Angriffe[i])
		}
	}
}
