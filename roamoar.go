package main

import (
    "log"
    "./models/being"
    "./models/location"
    "./travel"
)

func main() {
    b := being.NewToon("Purto")
    log.Print("b: ", b)
    b.SetBaseSpeed(1)
    l := location.New("Blarney", 50, 130, 5, 5)
    travel.MoveToward(b, l, 500)
}
