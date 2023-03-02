package playermap

import "fmt"

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

