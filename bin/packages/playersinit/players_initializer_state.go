package playersinit

func (players Players) IsOnOfThePlayersDefeated() bool {
    return players.FirstPlayer.PlayerMap.IsDefeated() || players.SecondPlayer.PlayerMap.IsDefeated()
}
