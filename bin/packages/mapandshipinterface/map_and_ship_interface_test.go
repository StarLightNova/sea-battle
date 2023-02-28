package mapandshipinterface

import (
  "testing"
  "fmt"
)

func TestNewMapShip(t *testing.T) {
  newMapShip := NewMapAndShipInitializer()

  fmt.Println(newMapShip)
}
