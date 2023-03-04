package mapandship

import (
	"fmt"
	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
)

type MapAndShip struct {
  ships []ShipQuantity;
  playerMap playermap.PlayerMap;
}

func New() MapAndShip {
  pMap, _ := playermap.New()

  return MapAndShip{standardShipsQuantity, *pMap}
}

func (masi MapAndShip) String() string {
  ans := ""

  for _, ship := range masi.ships {
    ans += fmt.Sprintf("%s\nAmount: %d\n", ship.ship, ship.quantity)
  }

  ans += fmt.Sprint("The map: \n", masi.playerMap)

  return ans
}
