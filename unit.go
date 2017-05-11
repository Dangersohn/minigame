package main

//Unit hat enthält eigenschaften der Einheiten
type Unit struct {
	Name     string
	Leben    int
	Armor    int
	Schaden  int
	AmLeben  bool
	Angriffe []string
}

//NewUnit legt einen neue Einheit an
func NewUnit(name string, leben, armor, schaden int, amleben bool, attacke, attacke2, attack3, attack4 string) *Unit {
	p := &Unit{
		Name:    name,
		Leben:   leben,
		Armor:   armor,
		Schaden: schaden,
		AmLeben: amleben,
		Angriffe: []string{
			attacke,
			attacke2,
			attack3,
			attack4,
		},
	}
	return p
}

//SetLeben setzt das leben
func (p *Unit) SetLeben(dmg int) {
	p.Leben = p.Leben - (dmg - p.Armor)
	p.CheckLife()
}

//CheckLife guckt ob dfas leben ueber 0 ist
func (p *Unit) CheckLife() bool {
	if p.Leben <= 0 {
		p.AmLeben = false
	}
	return p.AmLeben
}
