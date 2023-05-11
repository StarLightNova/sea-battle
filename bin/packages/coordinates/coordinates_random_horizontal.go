package coordinates

import (
    "math/rand"
    "github.com/StarLightNova/sea-battle/bin/packages/ship"
	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
)

func randomHorizontalCoordinates(ship ship.Ship) Coordinates {
    randomRow, _ := randomRow()
    colStart, colEnd := randomStartEndColumn(ship.CellsToOccupy)

    return Coordinates{
        StartRow: randomRow,
        EndRow: randomRow,
        StartColumn: colStart,
        EndColumn: colEnd,
    }
}

func randomRow() (string, int) {
    randomRowIndex := rand.Intn(len(playermap.GetLetterCoordinates()))

    return playermap.GetLetterCoordinates()[randomRowIndex], randomRowIndex
}

func randomStartEndRow(cellOccupation ship.ShipOccupy) (start, end string) {
    start, index := randomRow()

    if index + int(cellOccupation) > 10 {
        end = start
        start = playermap.GetLetterCoordinates()[abs(index - int(cellOccupation) + 1)]
    } else {
        end = playermap.GetLetterCoordinates()[index + int(cellOccupation) - 1]
    }

    return
}
