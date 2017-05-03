package unit

//Unit hat enthält eigenschaften der Einheiten
type Unit struct {
	Leben   int
	Armor   int
	Schaden int
	AmLeben bool
}

//SetLeben setzt das leben
func (p *Unit) SetLeben(dmg int) {
	//fmt.Println("in Set leben", dmg)
	p.Leben = p.Leben - dmg
	p.CheckLife()
}

//CheckLife guckt ob dfas leben ueber 0 ist
func (p *Unit) CheckLife() bool {
	if p.Leben <= 0 {
		p.AmLeben = false
	}
	return p.AmLeben
}
