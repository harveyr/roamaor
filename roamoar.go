package main

import (
    "fmt"
    "./models/being"
    "./models/location"
    "./travel"
)

func TickBeing(b *being.Being, time int) {
	fmt.Printf("--- Ticking %s ---\n", b)
	if b.Destination != nil {
	    travel.MoveToward(b, b.Destination, time)
	} else {
		fmt.Printf("%s has no Destination\n", b)
	}
	fmt.Println("")
}

func main() {
    b := being.NewToon("Purto")
    b.SetBaseSpeed(1)
    l := location.New("Blarney", 50, 130, 5, 5)
    TickBeing(b, 60)
    b.Destination = l.Start
    TickBeing(b, 60)
}
