package mapandship

import (
	"fmt"
	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
)

type MapAndShip struct {
  Ships []ShipQuantity;
  PlayerMap playermap.PlayerMap;
}

func New() MapAndShip {
  pMap, _ := playermap.New()

  return MapAndShip{standardShipsQuantity, *pMap}
}

func (masi MapAndShip) String() string {
  ans := ""

  for _, ship := range masi.Ships {
    ans += fmt.Sprintf("%s\nAmount: %d\n", ship.Ship, ship.quantity)
  }

  ans += fmt.Sprint("The map: \n", masi.PlayerMap)

  return ans
}
