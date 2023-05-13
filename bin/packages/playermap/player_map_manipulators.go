package playermap

func (playerMap PlayerMap) GetDamage(row string, column int) {
    if playerMap.GetCell(row, column) == "S" {
        playerMap.theMap[row][column - 1] = "X"
        playerMap.ShadowMap[row][column - 1] = "X"
    } else {
        playerMap.theMap[row][column - 1] = "M"
        playerMap.ShadowMap[row][column - 1] = "M"
    }
}

func (playerMap PlayerMap) PlaceUnit(row string, column int) {
    playerMap.theMap[row][column - 1] = "S"
}

func (playerMap PlayerMap) GetCell(row string, column int) string {
    return playerMap.theMap[row][column - 1]
}

func (playerMap PlayerMap) IsDefeated() bool {
    // If non "S" remains on the map, it is a lost case.
    for i := range playerMap.theMap {
        for _, jVal := range playerMap.theMap[i] {
            if jVal == "S" {
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

