package mapandship

import (
  "github.com/StarLightNova/sea-battle/bin/packages/ship"
  "github.com/StarLightNova/sea-battle/bin/packages/coordinates"
)

// Standard ships and quantity in the game.
type ShipQuantity struct {
  ship *ship.Ship;
  quantity int;
}

func newShipByClass(class ship.ShipClass) *ship.Ship {
  tempShip, _ := ship.New(class)
  return tempShip
}

func newShipAndQuantity(class ship.ShipClass, quantity int) ShipQuantity {
  return ShipQuantity{
    ship: newShipByClass(class),
    quantity: quantity,
  }
}

var standardShipsQuantity = []ShipQuantity{
  newShipAndQuantity(ship.Carrier, 1),
  newShipAndQuantity(ship.Battleship, 1),
  newShipAndQuantity(ship.Destroyer, 1),
  newShipAndQuantity(ship.Submarine, 2),
  newShipAndQuantity(ship.PatrolBoat, 2),
}

type Axis string

const (
  Horizontal = "horizontal"
  Vertical = "vertical"
  X = Horizontal
  Y = Vertical
)

func newCoordinate(strCoor string) *coordinates.Coordinates {
  coor, _ := coordinates.New(strCoor)
  return coor
}

// Some default coordinates to place the ships
var dummyCoordinates = []coordinates.Coordinates{
  // Carreir
  *newCoordinate("E10:J10"),
  // Battleship
  *newCoordinate("g1:j1"),
  // Destroyer
  *newCoordinate("f3:H3"),
  // Submarine #1
  *newCoordinate("J6:j8"),
  // Submarine #2
  *newCoordinate("d1:d3"),
  // Patrol Boat #1
  *newCoordinate("a1:b1"),
  // Patron Boat #2
  *newCoordinate("f6:f7"),
}
