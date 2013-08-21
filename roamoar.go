package main

import (
    "log"
    "./models/toon"
    "./models/location"
    "./travel"
)

func main() {
    t := new(toon.Toon)
    t.SetName("Purto")
    log.Print("t: ", t)
    t.SetBaseSpeed(1)
    l := new(location.Location)
    l.SetBounds(50, 130, 5, 5)
    travel.MoveToward(t, l, 500)
}
