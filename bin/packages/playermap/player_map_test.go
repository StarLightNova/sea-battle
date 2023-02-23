package playermap

import (
	"testing"
)

func TestPlayerMapCreated(t *testing.T) {
  _, error := NewPlayerMap()

  if error != nil {
    t.Fatal("Constructor FAILED! Got a bad PlayerMap object.")
  }
}

func TestPlayerMapSizings(t *testing.T) {
  playerMap, _ := NewPlayerMap()

  height := len(playerMap.theMap)
  width := len(playerMap.theMap["A"])

  if height != width || height != 10 || width != 10 {
    t.Fatal("The sizes of the map are different and not 10.")
  }
}

func TestPlayerMapConstructorReturnsEmptyMap(t *testing.T) {
  playerMap, _ := NewPlayerMap()

  for _, arrayOfLetters := range playerMap.theMap {
    for _, letter := range arrayOfLetters {
      if letter != "E" {
        t.Fatal("Brand new player map is *NOT* fully empty.")
      }
    }
  }
}

func TestMapDamageMethod(t *testing.T) {
  playerMap, _ := NewPlayerMap()

  playerMap.GetDamage("A", 1)

  if playerMap.theMap["A"][1] != "X" {
    t.Fatal("The damage did not set by the method GetDamage.")
  }
}
