package travel

import (
	"log"
	"math"
    "../models/toon"
    "../models/location"
)


func Distance(x1 float64, y1 float64, x2 float64, y2 float64) (d float64) {
	xDiff := math.Abs(x2 - x1)
	yDiff := math.Abs(y2 - y1)
	d = math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2))
	return
}

func MoveToward(t *toon.Toon, l *location.Location, time int) {
    log.Print("Moving")
    tX := float64(t.X)
    tY := float64(t.Y)
    lX := float64(l.X)
    lY := float64(l.Y)
    // slope := float32(l.Y - t.Y) / float32(l.X - t.X)
    potentialDistance := t.Speed() * float64(time) / 5
	potentialDistance = math.Min(1, potentialDistance)
    log.Print("potentialDistance ", potentialDistance)

    distanceToLocation := Distance(tX, tY, lX, lY)
    log.Print("distanceToLocation: ", distanceToLocation)

    var xMove, yMove float64 = 0, 0
    xDiff := (lX - tX)
    yDiff := (lY - tY)

    if xDiff == 0 && yDiff == 0 {
    	// We're there!
    	return
    }

    if xDiff == 0 {
    	// All vertical movement
    	yMove = potentialDistance
    } else if yDiff == 0 {
    	xMove = potentialDistance
    } else {
    	totalDiff := xDiff + yDiff
    	yMove = potentialDistance * yDiff / totalDiff
    	xMove = potentialDistance - yMove
    }
	log.Print("yMove: ", yMove)
	log.Print("xMove: ", xMove)
}
