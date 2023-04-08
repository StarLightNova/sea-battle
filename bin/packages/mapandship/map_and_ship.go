package mapandship

import (
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
    return masi.PlayerMap.String()
}

func (masi MapAndShip) ShadowString() string {
    return masi.PlayerMap.ShadowString()
}
