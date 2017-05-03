package userinterface

import (
	"fmt"
	"minigame/unit"
)

//Menu gibt den startus der Einheiten aus und die Atacken
func Menu(you unit.Unit, enemy unit.Unit) {
	fmt.Printf("Life: %d \t Enemy Life: %d\n", you.Leben, enemy.Leben)
}
