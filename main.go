package main

import (
	"fmt"
	"minigame/fight"
	"minigame/unit"
)

func menu(input int) {
	fmt.Println("tes")
}

func main() {
	//legt units an
	karl := &unit.Unit{Leben: 100, Armor: 5, Schaden: 20, AmLeben: true}
	dude := &unit.Unit{Leben: 100, Armor: 2, Schaden: 10, AmLeben: true}

	fight.Combat(karl, dude)
}
