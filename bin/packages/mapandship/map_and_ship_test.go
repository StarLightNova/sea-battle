package mapandship

import (
	"fmt"
	"testing"

	"github.com/StarLightNova/sea-battle/bin/packages/coordinates"
)

func TestMapShipUnitsPlacement(t *testing.T) {
  masi := New()
  coor, _ := coordinates.New("a1:b1")

  masi.placeShip(*coor)

  if masi.playerMap.GetCell("A", 1) != "S" && masi.playerMap.GetCell("B", 1) != "S" {
    t.Fatal("The ship did not placed on the passed coordinates.")
  }
}

func TestDummyCoordinatesPlacemen(t *testing.T) {
  masi := New()

  for _, coordinate := range dummyCoordinates {
    masi.placeShip(coordinate)
  }

  fmt.Println(masi)
}
