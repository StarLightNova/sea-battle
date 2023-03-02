package mapandshipinterface

import (
	"sort"

	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
)

// func (masi MapAndShipInterface) placeShipsOnTheMap() {

// }

// TODO:: Create a struct to handle the parameters: row, column
func (masi MapAndShipInterface) placeShip(startRow, endRow string, startColumn, endColumn int) {
  allRowKeys := playermap.GetLetterCoordinates();

  startRowIndex := sort.SearchStrings(allRowKeys[:], startRow)
  endRowIndex := sort.SearchStrings(allRowKeys[:], endRow)

  // Horizontal placing
  if startRow == endRow {
    for startColumn <= endColumn {
      masi.playerMap.PlaceUnit(startRow, startColumn)

      startColumn++
    }
  } else if startColumn == endColumn {
    for startRowIndex <= endRowIndex {
      masi.playerMap.PlaceUnit(allRowKeys[startRowIndex], startColumn)

      startRowIndex++
    }
  }
}
