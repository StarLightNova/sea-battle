package mapandship

import (
	"testing"
	"github.com/StarLightNova/sea-battle/bin/packages/coordinates"
  "fmt"
)

func TestMapShipUnitsPlacement(t *testing.T) {
  masi := New()
  coor, _ := coordinates.New("a1:b1")

  masi.placeShip(*coor)

  if masi.PlayerMap.GetCell("A", 1) != "S" && masi.PlayerMap.GetCell("B", 1) != "S" {
    t.Fatal("The ship did not placed on the passed coordinates.")
  }
}

func TestDummyCoordinatesPlacemen(t *testing.T) {
  masi := New()

  for _, coordinate := range dummyCoordinates {
    masi.placeShip(coordinate)
  }

  dummyCoor := dummyCoordinates[2]

  if masi.PlayerMap.GetCell(dummyCoor.StartRow, dummyCoor.StartColumn) != "S" {
    t.Fatal("The dummy coordinate ship placement is incorrect.")
  }
}

func TestRandomCoordinatesForShips(t *testing.T) {
  masi := New()
  
  for _, ship := range masi.Ships {
    randomCoordinates := coordinates.RandomCoordinatesFor(*ship.Ship)
    
    masi.placeShip(randomCoordinates)
  }

  fmt.Println(masi)
}
