package main

import (
	"fmt"
)

//Menu gibt den startus der Einheiten aus und die Atacken
func Menu(you *Unit, enemy *Unit) {
	fmt.Printf("Life: %d \t Enemy Life: %d\n", you.Leben, enemy.Leben)
}
