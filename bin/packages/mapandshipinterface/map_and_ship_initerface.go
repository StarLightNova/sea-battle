package mapandshipinterface

import (
	"fmt"
	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
)

type MapAndShipInterface struct {
  ships []ShipQuantity;
  playerMap playermap.PlayerMap;
}

func NewMapAndShipInitializer() MapAndShipInterface {
  pMap, _ := playermap.NewPlayerMap()

  return MapAndShipInterface{standardShipsQuantity, *pMap}
}

func (masi MapAndShipInterface) String() string {
  ans := ""

  for _, ship := range masi.ships {
    ans += fmt.Sprintf("%s\nAmount: %d\n", ship.ship, ship.quantity)
  }

  ans += fmt.Sprint("The map: \n", masi.playerMap)

  return ans
}
