package mapandship

import (
	"sort"
	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
  "github.com/StarLightNova/sea-battle/bin/packages/coordinates"
)

func (masi MapAndShip) placeShip(coor coordinates.Coordinates) {
  // Horizontal placing
  if coor.StartRow == coor.EndRow {
    masi.forPlacer(coor, coor.StartColumn, Horizontal)
  } else if coor.StartColumn == coor.EndColumn {
    masi.forPlacer(coor, 0, Vertical)
  }
}

func (masi MapAndShip) forPlacer(coor coordinates.Coordinates, incrementer int, axis Axis) {
  if axis == Vertical {
    allRowKeys := playermap.GetLetterCoordinates();
    startRowIndex := sort.SearchStrings(allRowKeys[:], coor.StartRow)
    endRowIndex := sort.SearchStrings(allRowKeys[:], coor.EndRow)

    for startRowIndex <= endRowIndex {
      masi.playerMap.PlaceUnit(allRowKeys[startRowIndex], coor.StartColumn)

      startRowIndex++
    }
  } else {
    for coor.StartColumn <= coor.EndColumn {
      masi.playerMap.PlaceUnit(coor.StartRow, incrementer)

      incrementer++
    }
  }
}
