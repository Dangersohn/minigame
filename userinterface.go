package main

import (
	"fmt"
)

//Menu gibt den startus der Einheiten aus und die Atacken
func Menu(you *Unit, enemy *Unit) {
	fmt.Printf("Life: %d \t Enemy Life: %d\n", you.Leben, enemy.Leben)
	Atkausgabe(you)
}

//Atkausgabe test blub
func Atkausgabe(you *Unit) {
	for i := 0; i < len(you.Angriffe); i++ {
		if you.Angriffe[i] != "" {
			fmt.Printf("%d) %s \n", i+1, you.Angriffe[i])
		}
	}
}

//Angriffwahl logig hinter dem angriff menu
func Angriffwahl(you *Unit, enemy *Unit) {
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("eingabe", input)
	switch input {
	case 1:
		SchlagNormal(you, enemy)
	case 2:
		SchlagSchwer(you, enemy)
	case 3:
		NinjaKick(you, enemy)
	case 4:
		Wurf(you, enemy)
	}
}
