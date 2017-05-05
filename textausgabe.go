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

func textausgabe() {

	f, err := os.Open("choices/choice1.1.txt")
	check(err)
	ff, err := f.Stat()
	fff := ff.Size()
	var i int64
	for i = 0; i < fff; i++ {
		b1 := make([]byte, 1)
		f.Read(b1)
		check(err)
		fmt.Printf("%s", string(b1))
		time.Sleep(60000000)
	}
}
