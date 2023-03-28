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

func randomVerticalCoordinates(ship ship.Ship) Coordinates {
  rand.Seed(time.Now().UnixNano())
  randomColumn := randomColumn()
  rowStart, rowEnd := randomStartEndRow(ship.CellsToOccupy)

  return Coordinates{
    StartRow: rowStart,
    EndRow: rowEnd,
    StartColumn: randomColumn,
    EndColumn: randomColumn,
  }
}

func randomHorizontalCoordinates(ship ship.Ship) Coordinates {
  rand.Seed(time.Now().UnixNano())
  randomRow, _ := randomRow()
  colStart, colEnd := randomStartEndColumn(ship.CellsToOccupy)

  return Coordinates{
    StartRow: randomRow,
    EndRow: randomRow,
    StartColumn: colStart,
    EndColumn: colEnd,
  }
}

func randomAxis() Axis {
  rand.Seed(time.Now().UnixNano())

  if rand.Float32() > 0.5 {
    return Vertical
  }

  return Horizontal
}

func randomRow() (string, int) {
  rand.Seed(time.Now().UnixNano())

  randomRowIndex := rand.Intn(len(playermap.GetLetterCoordinates()))

  return playermap.GetLetterCoordinates()[randomRowIndex], randomRowIndex
}

func randomColumn() int {
  rand.Seed(time.Now().UnixNano())

  return rand.Intn(10) + 1
}

func randomStartEndColumn(cellOccupation ship.ShipOccupy) (start, end int) {
  rand.Seed(time.Now().UnixNano())
  start = randomColumn()

  if start + int(cellOccupation) > 10 {
    end = start
    start = start - int(cellOccupation)
  } else {
    end = start + int(cellOccupation)
  }

  return
}

func randomStartEndRow(cellOccupation ship.ShipOccupy) (start, end string) {
  rand.Seed(time.Now().UnixNano())
  start, index := randomRow()

  if index + int(cellOccupation) > 10 {
    end = start
    start = playermap.GetLetterCoordinates()[index - int(cellOccupation)]
  } else {
    end = playermap.GetLetterCoordinates()[index + int(cellOccupation) - 1]
  }

  return
}
