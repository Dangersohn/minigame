package main

import (
	"fmt"
	"math/rand"
)

//Combat läuft so lange bis einer der betiligten tot ist
func Combat(you *Unit, enemy *Unit) {
	for enemy.AmLeben && you.AmLeben {
		if enemy.AmLeben && you.AmLeben {
			Menu(you, enemy)
			Angriffwahl(you, enemy)
		}

		if enemy.AmLeben && you.AmLeben {
			Randatk(enemy, you)
		}
	}
	if you.AmLeben {
		println("Rufus gewinnt den Kampf")
	}
}

//Randatk wählt einen zufälligen angriff
func Randatk(enemy *Unit, you *Unit) {
	zufall := rand.Intn(100)
	fmt.Println("Zufall", zufall)
	if zufall <= 25 {
		SchlagNormal(enemy, you)
	} else if zufall <= 50 {
		SchlagSchwer(enemy, you)
	} else if zufall <= 75 {
		NinjaKick(enemy, you)
	} else {
		Wurf(enemy, you)
	}
}
