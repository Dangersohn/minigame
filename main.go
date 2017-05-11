package main

func main() {
	//legt units an
	rufus := NewUnit("Rufus", 100, 5, 20, true, "Schlag Normal", "Schlag Schwer", "Ninja Kick", "Wurf")
	schnappschild := NewUnit("Schnappschildkröte", 100, 5, 20, true, "Schlag Normal", "Schlag Schwer", "", "")
	//Combat(karl, dude)

	textausgabe("intro.txt")

	enscheidungen := enscheidung("choice1.txt", "Kanalisation", "choice2.txt", "Gleitdrache")
	// erste enscheidung
	if enscheidungen == 1 {
		enscheidung("choice1.1.txt", "Beeilen", "choice1.2.txt", "Prüfen")
	}
	//zweite enscheidung
	if enscheidungen == 2 {
		enscheidung("choice2.1.txt", "Springen", "choice2.2.txt", "Hinablassen")
	}

	//Im Gebäude angekommen
	textausgabe("sequens.txt")

	enscheidung("choice3.1.txt", "Renne", "choice3.1.txt", "Seilschleuder")

	Combat(rufus, schnappschild)

	//nach dem Kampf
	enscheidung("choice4.1.txt", "Schacht", "choice4.2.txt", "Gang")

	textausgabe("herz.txt")

}
