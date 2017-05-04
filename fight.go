package main

//Combat läuft so lange bis einer der betiligten tot ist
func Combat(you *Unit, enemy *Unit) {
	for enemy.AmLeben {
		Menu(you, enemy)
		enemy.SetLeben(SchlagNormal(you))
		//fmt.Println("leben gegner", enemy.Leben)
	}
}
