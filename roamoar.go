package main

import (
    // "log"
    // "fmt"
    "time"
    "math/rand"
    "./domain"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	domain.InitDb("localhost", "roamaor")

	toons := domain.FetchAllToons()
	for _, toon := range toons {
		domain.TickBeing(&toon)
	}
	// time.Sleep(1 * time.Second)

    // toon.SetBaseSpeed(1)
    // l := NewLocation("Blarney", 50, 130, 5, 5)
    // TickBeing(toon, 60)
    // toon.Destination = l.Start
    // TickBeing(toon, 60)

    // w := NewWeapon(5)
    // fmt.Printf("%s", w)

    // mob := NewMob(1)
    // Fight(mob, toon)
}
