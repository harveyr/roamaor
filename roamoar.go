package main

import (
    "fmt"
    "time"
    "math/rand"
    "./models/being"
    "./models/location"
    "./models/weapon"
    "./travel"
    "./fight"
)

func TickBeing(b *being.Being, seconds int) {
	fmt.Printf("--- Ticking %s ---\n", b)
	fmt.Println("Since last tick: ", b.SinceLastTick())
	if b.Destination != nil {
	    travel.MoveToward(b, b.Destination, seconds)
	} else {
		fmt.Printf("%s has no Destination\n", b)
	}
	b.UpdateLastTick()
	time.Sleep(1 * time.Second)
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
    toon := being.NewToon("Purto")
    toon.SetBaseSpeed(1)
    l := location.New("Blarney", 50, 130, 5, 5)
    TickBeing(toon, 60)
    toon.Destination = l.Start
    TickBeing(toon, 60)

    w := weapon.NewWeapon(5)
    fmt.Printf("%s", w)

    mob := being.NewMob(1)
    fight.Fight(mob, toon)
}
