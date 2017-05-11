package main

import (
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func textausgabe(datei string) {
	fmt.Println("")
	datei = "choices/" + datei
	f, err := os.Open(datei)
	check(err)
	ff, err := f.Stat()
	fff := ff.Size()
	var i int64
	for i = 0; i < fff; i++ {
		b1 := make([]byte, 1)
		f.Read(b1)
		check(err)
		fmt.Printf("%s", string(b1))
		time.Sleep(60000)
	}
}

func enscheidung(wahl1, wahl1Ausgabe, wahl2, wahl2Ausgabe string) int {
	fmt.Printf("\n\t1) %s\n\t2) %s \n=> ", wahl1Ausgabe, wahl2Ausgabe)
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println(err)
	}
	if input == 1 {
		textausgabe(wahl1)
	}
	if input == 2 {
		textausgabe(wahl2)
	} else {
		fmt.Println("flasche eingabe")
	}
	return input

}
