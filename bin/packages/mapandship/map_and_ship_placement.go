package mapandship

import (
    "sort"
    "github.com/StarLightNova/sea-battle/bin/packages/coordinates"
    "github.com/StarLightNova/sea-battle/bin/packages/playermap"
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
            masi.PlayerMap.PlaceUnit(allRowKeys[startRowIndex], coor.StartColumn)

            startRowIndex++
        }
    } else {
        for incrementer <= coor.EndColumn {
            masi.PlayerMap.PlaceUnit(coor.StartRow, incrementer)

            incrementer++
        }
    }
}

func (masi MapAndShip) UniqPlacement() {
    for _, ship := range masi.Ships {
        shipAmount := ship.quantity

        for shipAmount > 0 {
            randomCoordinates := coordinates.RandomCoordinatesFor(*ship.Ship)

            for masi.isOverlapping(randomCoordinates) {
                randomCoordinates = coordinates.RandomCoordinatesFor(*ship.Ship)
            }

            masi.placeShip(randomCoordinates)

            shipAmount--
        }
    }
}

func (masi MapAndShip) isOverlapping(coor coordinates.Coordinates) bool {
    if (coor.StartRow == coor.EndRow) {
        for i := coor.StartColumn; i <= coor.EndColumn; i++ {
            if masi.PlayerMap.GetCell(coor.StartRow, i) == "S" {
                return true;
            }
        }
    } else {
        allRowKeys := playermap.GetLetterCoordinates();
        startRowIndex := sort.SearchStrings(allRowKeys[:], coor.StartRow)
        endRowIndex := sort.SearchStrings(allRowKeys[:], coor.EndRow)

        for startRowIndex <= endRowIndex {
            if masi.PlayerMap.GetCell(allRowKeys[startRowIndex], coor.StartColumn) == "S" {
                return true;
            }

            startRowIndex++
        }
    }

    return false
}

