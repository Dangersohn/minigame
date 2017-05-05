package main

//SchlagNormal berechnet den Schaden den er Angreifer machen wird
func SchlagNormal(angreifer *Unit, verteidiger *Unit) {
	verteidiger.SetLeben(angreifer.Schaden * 2)
}

//SchlagSchwer ist ein ungenauer angriff macht viel Schaden
func SchlagSchwer(angreifer *Unit, verteidiger *Unit) {
	verteidiger.SetLeben(angreifer.Schaden * 5)
}
