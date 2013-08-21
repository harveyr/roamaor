package travel

import (
	"log"
	"math"
	"fmt"
    "../models"
    "../models/being"
    "../models/location"
)


func Distance(x1 float64, y1 float64, x2 float64, y2 float64) (d float64) {
	xDiff := math.Abs(x2 - x1)
	yDiff := math.Abs(y2 - y1)
	d = math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2))
	return
}

func DistBetw(p1 location.Point, p2 location.Point) float64 {
	return Distance(float64(p1.X), float64(p1.Y), float64(p2.X), float64(p2.Y))
}

func MoveToward(b *being.Being, l *location.Location, time int) {
    log.Print(fmt.Sprintf("Moving %s toward %s", b, l))
    bX := float64(b.X)
    bY := float64(b.Y)
    lX := float64(l.X)
    lY := float64(l.Y)
    // log.Print(fmt.Sprintf("%s begins at distance %f", b, Distance()))
    DistBetw(b, l)

    potentialDistance := math.Min(1, b.Speed() * float64(time) / 5)

    var xMove, yMove float64 = 0, 0
    xDiff := (lX - bX)
    yDiff := (lY - bY)

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

    yMove = math.Min(yMove, math.Abs(lY - bY))
    xMove = math.Min(xMove, math.Abs(lX - bX))
	// log.Print("yMove: ", yMove)
	// log.Print("xMove: ", xMove)
	b.X += float32(xMove)
	b.Y += float32(yMove)
}
