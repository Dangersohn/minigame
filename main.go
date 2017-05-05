package main

func main() {
	//legt units an
	karl := NewUnit(100, 5, 20, true, "Schlag Normal", "Schlag Schwer", "Ninja Kick")
	dude := NewUnit(100, 5, 20, true, "Schlag Normal", "Schlag Schwer", "")
	Combat(karl, dude)
}
