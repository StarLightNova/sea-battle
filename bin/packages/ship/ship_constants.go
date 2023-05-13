package ship

type ShipClass string;
type ShipOccupy int;

const (
    Carrier ShipClass = "Carrier"
    Battleship ShipClass = "Battleship"
    Destroyer ShipClass = "Destroyer"
    Submarine ShipClass = "Submarine"
    PatrolBoat ShipClass = "Patrol Boat"
)

var shipsAndOccupyCells = map[ShipClass]ShipOccupy{
    Carrier: 5,
    Battleship: 4,
    Destroyer: 3,
    Submarine: 3,
    PatrolBoat: 2,
}
