package fight

import (
	"minigame/attacks"
	"minigame/unit"
	"minigame/userinterface"
)

//Combat läuft so lange bis einer der betiligten tot ist
func Combat(you unit.Unit, enemy unit.Unit) {
	for enemy.AmLeben {
		userinterface.Menu(you, enemy)
		enemy.SetLeben(attacks.SchlagNormal(you))
		//fmt.Println("leben gegner", enemy.Leben)
	}
}
