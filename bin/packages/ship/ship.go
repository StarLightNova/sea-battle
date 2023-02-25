package ship

import (
  "fmt"
  "errors"
)

type ShipClass string;
type ShipOccupy int;

const (
  Carrier ShipClass = "Carrier"
  Battleship ShipClass = "Battleship"
  Destroyer ShipClass = "Destroyer"
  Submarine ShipClass = "Submarine"
  PatrolBoat ShipClass = "Patrol Boat"
)

type Ship struct {
  class ShipClass;
  cellsToOccupy ShipOccupy;
}

func NewShip(class ShipClass) (*Ship, error) {
  shipCells := map[ShipClass]ShipOccupy {
    Carrier: 5,
    Battleship: 4,
    Destroyer: 3,
    Submarine: 3,
    PatrolBoat: 2,
  }

  _, isShipExists := shipCells[class]

  if !isShipExists {
    return nil, errors.New("The ship of this class does not exist.")
  }

  return &Ship {
    class: class,
    cellsToOccupy: shipCells[class],
  }, nil
}

func (ship Ship) String() string {
  return fmt.Sprintf("The ship: %s, occupies: %d", ship.class, ship.cellsToOccupy)
}
