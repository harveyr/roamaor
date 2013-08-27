package domain

import (
	"fmt"
	"math"
	"math.rand"
	"time"
)

var _ string = fmt.Sprintf("")
var _ time.Time = time.Now()

func EarnedLevel(b *Being) (level int) {
	// daysInGame := float64(time.Now().UTC().Sub(b.Created)) / float64(time.Hour * 24)
	winsRequired := int(math.Pow(float64(b.Level), 2)) + b.Level
	if b.FightsWon > winsRequired {
		level = b.Level + 1
	}
	return
}

func ApplyProgress(b *Being) {
	earnedLevel := EarnedLevel(b)
	if earnedLevel > b.Level {
		diff := earnedLevel - b.Level
		b.Level = earnedLevel
		b.MaxHp += 20 + earnedLevel + rand.Intn(earnedLevel)
		b.Save()
	}
}
