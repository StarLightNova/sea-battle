package playermap

import "fmt"

const Width = 10
const Height = Width

type PlayerMap struct {
  theMap map[string][]string
}

func NewPlayerMap() (*PlayerMap, error) {
  pureMap := make(map[string][]string)

  fillPureMap(pureMap)

  return &PlayerMap {
    theMap: pureMap,
  }, nil
}

func (playerMap PlayerMap) GetDamage(row string, column int) {
  playerMap.theMap[row][column - 1] = "X"
}

func (playerMap PlayerMap) GetCell(row string, column int) string {
  return playerMap.theMap[row][column - 1]
}

func (playerMap PlayerMap) String() string {
  returnString := fmt.Sprintf("     1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10\n")
  returnString += "===========================================\n"

  for key, arrayOfElements := range playerMap.theMap {
    returnString += fmt.Sprintf(" %s )", key)

    for _, letter := range arrayOfElements {
      returnString += fmt.Sprintf(" %s |", letter)
    }

    returnString += "\n-------------------------------------------\n"
  }

  return returnString
}


func GetLetterCoordinates() [Width]string {
  return [Width]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
}

/* [----PRIVATE----] */

func fillPureMap(pureMap map[string][]string) {
   /*
   * Fill the playerMap variable by keys from 
   * `getLetterCoordinates` and by numbers till `Width`.
   * All the cells should be 'empty' by letter 'E'.
  */

  for _, key := range GetLetterCoordinates() {
    pureMap[key] = []string{"E", "E", "E", "E", "E", "E", "E", "E", "E", "E"}
  }
}

