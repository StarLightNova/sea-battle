package playersinit

func (players Players) PutShips() {
    players.FirstPlayer.UniqPlacement()
    players.SecondPlayer.UniqPlacement()
}

func (players Players) IsOnOfThePlayersDefeated() bool {
    return players.FirstPlayer.PlayerMap.IsDefeated() || players.SecondPlayer.PlayerMap.IsDefeated()
}
