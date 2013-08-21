package travel

import (
	"math"
	"fmt"
    "../models/being"
    "../models/location"
)


func Distance(x1 float64, y1 float64, x2 float64, y2 float64) (d float64) {
	xDiff := math.Abs(x2 - x1)
	yDiff := math.Abs(y2 - y1)
	d = math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2))
	return
}

func DistBetw(p1 *location.Point, p2 *location.Point) float64 {
	return Distance(float64(p1.X), float64(p1.Y), float64(p2.X), float64(p2.Y))
}

func MoveToward(b *being.Being, p *location.Point, time int) {
    fmt.Printf("Moving %s toward %s\n", b, p)
    bX := float64(b.Location.X)
    bY := float64(b.Location.Y)
    pX := float64(p.X)
    pY := float64(p.Y)
    fmt.Printf("... %s begins at distance %f\n", b, DistBetw(b.Location, p))

    potentialDistance := math.Min(1, b.Speed() * float64(time) / 5)

    var xMove, yMove float64 = 0, 0
    xDiff := (pX - bX)
    yDiff := (pY - bY)

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

    yMove = math.Min(yMove, math.Abs(pY - bY))
    xMove = math.Min(xMove, math.Abs(pX - bX))
	b.Location.X += float32(xMove)
	b.Location.Y += float32(yMove)
    fmt.Printf("... %s ends at distance %f\n", b, DistBetw(b.Location, p))
}
