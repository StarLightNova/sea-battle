package playermap

import (
    "testing"
)

func TestPlayerMapCreated(t *testing.T) {
    _, error := New()

    if error != nil {
        t.Fatal("Constructor FAILED! Got a bad PlayerMap object.")
    }
}

func TestPlayerMapSizings(t *testing.T) {
    playerMap, _ := New()

    height := len(playerMap.TheMap)
    width := len(playerMap.TheMap["A"])

    if height != width || height != 10 || width != 10 {
        t.Fatal("The sizes of the map are different and not 10.")
    }
}

func TestPlayerMapConstructorReturnsEmptyMap(t *testing.T) {
    playerMap, _ := New()

    for _, arrayOfLetters := range playerMap.TheMap {
        for _, letter := range arrayOfLetters {
            if letter != " " {
                t.Fatal("Brand new player map is *NOT* fully empty.")
            }
        }
    }
}

func TestGetCellOfTheMapMethod(t *testing.T) {
    playerMap, _ := New()

    playerMap.GetCell("A", 1)

    if playerMap.GetCell("A", 1) != " " {
        t.Fatal("The GetCell() function returned non-empty cell.")
    }

}

func TestMapDamageMethod(t *testing.T) {
    playerMap, _ := New()

    playerMap.PlaceUnit("A", 1)
    playerMap.GetDamage("A", 1)

    if playerMap.GetCell("A", 1) != "X" {
        t.Fatal("The damage did not set by the method GetDamage.")
    }
}

func TestMapPlaceUnit(t *testing.T) {
    playerMap, _ := New()

    playerMap.PlaceUnit("B", 3)

    if playerMap.GetCell("B", 3) != "S" {
        t.Fatal("The unit's cells not placed.")
    }
}

func TestIsDeafeated(t *testing.T) {
    playerMap, _ := New()

    playerMap.PlaceUnit("B", 3)
    playerMap.GetDamage("B", 3)

    playerMap.PlaceUnit("G", 5)
    playerMap.GetDamage("G", 5)

    if !playerMap.IsDefeated() {
        t.Fatal("The player is not defeated. But should be.")
    }
}
