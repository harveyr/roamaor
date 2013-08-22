package main

import (
    "fmt"
    "time"
    "math/rand"
    . "./domain"
)

func TickBeing(b *Being, seconds int) {
	fmt.Printf("--- Ticking %s ---\n", b)
	fmt.Println("Since last tick: ", b.SinceLastTick())
	if b.Destination != nil {
	    MoveToward(b, b.Destination, seconds)
	} else {
		fmt.Printf("%s has no Destination\n", b)
	}
	b.UpdateLastTick()
	time.Sleep(1 * time.Second)
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
    toon := NewToon("Purto")
    toon.SetBaseSpeed(1)
    l := NewLocation("Blarney", 50, 130, 5, 5)
    TickBeing(toon, 60)
    toon.Destination = l.Start
    TickBeing(toon, 60)

    w := NewWeapon(5)
    fmt.Printf("%s", w)

    mob := NewMob(1)
    Fight(mob, toon)
}
