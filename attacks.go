package main

import "fmt"

//SchlagNormal berechnet den Schaden den er Angreifer machen wird
func SchlagNormal(angreifer *Unit, verteidiger *Unit) {
	fmt.Printf("%s greift %s mit Schlag Normal an\n", angreifer.Name, verteidiger.Name)
	verteidiger.SetLeben(angreifer.Schaden * 1)
}

//SchlagSchwer ist ein ungenauer angriff macht viel Schaden
func SchlagSchwer(angreifer *Unit, verteidiger *Unit) {
	fmt.Printf("%s greift %s mit Schlag Schwer an\n", angreifer.Name, verteidiger.Name)
	verteidiger.SetLeben(angreifer.Schaden * 3)
}

//NinjaKick schneller genauer angriff
func NinjaKick(angreifer *Unit, verteidiger *Unit) {
	fmt.Printf("%s greift %s mit Ninjakick an\n", angreifer.Name, verteidiger.Name)
	verteidiger.SetLeben(angreifer.Schaden * 2)
}

//Wurf ist ein wurfangriff
func Wurf(angreifer *Unit, verteidiger *Unit) {
	fmt.Printf("%s greift %s mit Wurf an\n", angreifer.Name, verteidiger.Name)
	verteidiger.SetLeben(angreifer.Schaden * 1)
}
