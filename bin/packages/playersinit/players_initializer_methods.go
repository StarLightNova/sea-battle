package playersinit

func (players Players) PutShips() {
  players.FirstPlayer.UniqPlacement()
  players.SecondPlayer.UniqPlacement()
}
