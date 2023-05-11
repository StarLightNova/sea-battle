package coordinates

import (
	"math/rand"
	"time"
	"github.com/StarLightNova/sea-battle/bin/packages/ship"
	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
)

func RandomCoordinatesFor(ship ship.Ship) Coordinates {
    rand.Seed(time.Now().UnixNano())

    if (randomAxis() == Horizontal) {
        return randomHorizontalCoordinates(ship)
    }

    return randomVerticalCoordinates(ship)
}

func RandomHalfCoordinates() (row string, column int) {
    rand.Seed(time.Now().UnixNano())

    row = playermap.GetLetterCoordinates()[rand.Intn(len(playermap.GetLetterCoordinates()))]
    column = rand.Intn(10) + 1

    return
}

func randomAxis() Axis {
    if rand.Float32() > 0.5 {
        return Vertical
    }

    return Horizontal
}

func abs(number int) int {
    if number >= 0 {
        return number
    }

    return -number
}
