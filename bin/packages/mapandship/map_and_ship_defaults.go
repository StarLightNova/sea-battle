package mapandship

import "github.com/StarLightNova/sea-battle/bin/packages/ship"

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

