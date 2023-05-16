package playermap

func (playerMap PlayerMap) GetDamage(row string, column int) {
    if playerMap.GetCell(row, column) == SHIP {
        playerMap.TheMap[row][column - 1] = HIT
        playerMap.ShadowMap[row][column - 1] = HIT
    } else if playerMap.GetCell(row, column) == EMPTY {
        playerMap.TheMap[row][column - 1] = MISS
        playerMap.ShadowMap[row][column - 1] = MISS
    }
}

func (playerMap PlayerMap) PlaceUnit(row string, column int) {
    playerMap.TheMap[row][column - 1] = SHIP
}

func (playerMap PlayerMap) GetCell(row string, column int) string {
    return playerMap.TheMap[row][column - 1]
}

func (playerMap PlayerMap) IsDefeated() bool {
    // If non "S" remains on the map, it is a lost case.
    for i := range playerMap.TheMap {
        for _, jVal := range playerMap.TheMap[i] {
            if jVal == SHIP {
                return false
            }
        }
    }

    return true
}

func fillPureMap(pureMap map[string][]string) {
    // All the cells should be 'empty' by letter ' ' (space).
    for _, key := range GetLetterCoordinates() {
        pureMap[key] = emptyRow()
    }
}

