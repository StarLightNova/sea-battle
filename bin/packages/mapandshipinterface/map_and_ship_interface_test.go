package mapandshipinterface

import (
  "testing"
)

func TestMapShipUnitsPlacement(t *testing.T) {
  masi := NewMapAndShipInitializer()

  masi.placeShip("A", "B", 1, 1)

  if masi.playerMap.GetCell("A", 1) != "S" && masi.playerMap.GetCell("B", 1) != "S" {
    t.Fatal("The ship did not placed on the passed coordinates.")
  }
}
