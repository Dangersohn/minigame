package main

import (
	"fmt"
)

//Unit hat enthält eigenschaften der Einheiten
type Unit struct {
	Leben   int
	Armor   int
	Schaden int
	AmLeben bool
}

//SchlagNormal ist die einfache angriffsfunktion
func (p *Unit) SchlagNormal() int {
	return p.Schaden * 2
}

//Hit rechnet den schaden gegen die Armor
func (p *Unit) Hit(dmg int) {
	p.Leben = p.Leben - dmg
}

func main() {
	//legt units an
	karl := Unit{Leben: 100, Armor: 5, Schaden: 20, AmLeben: true}
	dude := Unit{Leben: 100, Armor: 2, Schaden: 10, AmLeben: true}
	fmt.Println(karl.Leben, dude.Leben)
	dude.Hit(karl.SchlagNormal())

	fmt.Println(dude.Leben)

}
