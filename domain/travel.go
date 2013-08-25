package domain

import (
	"math"
	"log"
)

func Distance(x1 float64, y1 float64, x2 float64, y2 float64) (d float64) {
	xDiff := math.Abs(x2 - x1)
	yDiff := math.Abs(y2 - y1)
	d = math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2))
	return
}

func Roam(b *Being, time float64) {
    potentialDistance := math.Min(1, b.Speed() * time / 5)

    var xMove, yMove float64 = 0, 0
    xDiff := (b.DestX - b.LocX)
    yDiff := (b.DestY - b.LocY)

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

    yMove = math.Min(yMove, math.Abs(xDiff))
    xMove = math.Min(xMove, math.Abs(yDiff))
	b.LocX += xMove
	b.LocY += yMove
	b.Save()
	log.Printf("\tlocation: {%f, %f}", b.LocX, b.LocY)
}


func UpdateLocationsVisited(b *Being) {
	
}
