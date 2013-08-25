package main

import (
    "log"
    "flag"
    "time"
    "math/rand"
    "./domain"
)

func main() {
	log.Print("Roamaor starting up ...")

	var ticks int
	flag.IntVar(&ticks, "t", 1, "number of ticks")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	log.Print("Initializing database...")
	domain.InitDb("localhost", "roamaor")
	log.Print("Initializing world...")
	domain.InitWorldLocations()

	toons := domain.FetchAllToons()
	for i := 0; i < ticks; i++ {
		log.Printf("Starting tick %d / %d", i, ticks)
		for _, toon := range toons {
			toon.Reload()
			domain.TickBeing(&toon)
		}
		time.Sleep(1 * time.Second)
	}
}
