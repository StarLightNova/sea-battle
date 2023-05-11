package coordinates

import (
    "math/rand"
    "github.com/StarLightNova/sea-battle/bin/packages/ship"
)

func randomVerticalCoordinates(ship ship.Ship) Coordinates {
    randomColumn := randomColumn()
    rowStart, rowEnd := randomStartEndRow(ship.CellsToOccupy)

    return Coordinates{
        StartRow: rowStart,
        EndRow: rowEnd,
        StartColumn: randomColumn,
        EndColumn: randomColumn,
    }
}

func randomColumn() int {
    return rand.Intn(10) + 1
}

func randomStartEndColumn(cellOccupation ship.ShipOccupy) (start, end int) {
    start = randomColumn()

    if start + int(cellOccupation) > 10 {
        end = start
        start = abs(start - int(cellOccupation) + 1)
    } else {
        end = start + int(cellOccupation) - 1
    }

    return
}

