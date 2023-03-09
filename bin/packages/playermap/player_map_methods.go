package playermap

const Width = 10
const Height = Width

func (playerMap PlayerMap) GetDamage(row string, column int) {
  playerMap.theMap[row][column - 1] = "X"
}

func (playerMap PlayerMap) PlaceUnit(row string, column int) {
  playerMap.theMap[row][column - 1] = "S"
}

func (playerMap PlayerMap) GetCell(row string, column int) string {
  return playerMap.theMap[row][column - 1]
}

func GetLetterCoordinates() [Width]string {
  return [Width]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
}

func fillPureMap(pureMap map[string][]string) {
   /*
   * Fill the playerMap variable by keys from 
   * `getLetterCoordinates` and by numbers till `Width`.
   * All the cells should be 'empty' by letter ' ' (space).
  */

  for _, key := range GetLetterCoordinates() {
    pureMap[key] = []string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "}
  }
}


