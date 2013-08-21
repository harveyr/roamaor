package main

import (
    "log"
    "./models/being"
    "./models/location"
    "./travel"
)

func main() {
    b := new(being.Being)
    b.SetName("Purto")
    log.Print("b: ", b)
    b.SetBaseSpeed(1)
    l := new(location.Location)
    l.SetBounds(50, 130, 5, 5)
    travel.MoveToward(b, l, 500)
}
