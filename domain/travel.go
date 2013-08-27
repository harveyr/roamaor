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
    potentialDistance := b.Speed() * time / 60

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
    	totalDiff := math.Abs(xDiff) + math.Abs(yDiff)
    	yMove = potentialDistance * math.Abs(yDiff) / totalDiff
    	xMove = potentialDistance - yMove
    }

    xMove = math.Min(xMove, math.Abs(xDiff))
    yMove = math.Min(yMove, math.Abs(yDiff))

    if xDiff < 0 {
    	xMove *= -1
    }
    if yDiff < 0 {
    	yMove *= -1
    }

	b.LocX += xMove
	b.LocY += yMove

	b.Save()
	
	UpdateLocationsVisited(b)
}

func UpdateLocationsVisited(b *Being) {
	locs := FetchLocationsAt(b.LocX, b.LocY)
	if len(locs) == 0 {
		return
	}
	for _, loc := range locs {
		inVisited := false
		for _, visitedLoc := range b.LocationsVisited {
			if visitedLoc == loc.Id {
				inVisited = true
				break
			}
		}
		if !inVisited {
			b.LocationsVisited = append(b.LocationsVisited, loc.Id)
			item := NewLogItem(b, LOG_LOCATION_DISCOVERY)
			item.SetAttr("locationName", loc.Name)
			item.Save()
		}
	}
	b.Save()
}
