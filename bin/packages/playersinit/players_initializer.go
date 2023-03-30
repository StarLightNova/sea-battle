package playersinit

import (
  	"github.com/StarLightNova/sea-battle/bin/packages/mapandship"
)

type Players struct {
  FirstPlayer mapandship.MapAndShip;
  SecondPlayer mapandship.MapAndShip;
}

func New() (Players, error) {
  return Players{
    FirstPlayer: mapandship.New(),
    SecondPlayer: mapandship.New(),
  }, nil
}
