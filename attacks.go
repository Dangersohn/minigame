package main

//SchlagNormal berechnet den Schaden den er Angreifer machen wird
func SchlagNormal(p *Unit) int {
	return p.Schaden * 2
}
