package attacks

import (
	"minigame/unit"
)

//SchlagNormal berechnet den Schaden den er Angreifer machen wird
func SchlagNormal(p unit.Unit) int {
	return p.Schaden * 2
}
