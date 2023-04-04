package playersinit

import (
    "fmt"
    "testing"
)

func TestPlayerConstructor(t *testing.T) {
    _, error := New()

    if error != nil {
        t.Fatal("The players wrapper (initializer) didn't create an instance.")
    }
}

func TestPlayersMapsandShipPlacements(t *testing.T) {
    players, error := New()

    players.PutShips()

    fmt.Println("The first Player:", players.FirstPlayer)
    fmt.Println()
    fmt.Println("Second Player", players.SecondPlayer)


    if error != nil {
        t.Fatal("The players map are incorrect.")
    }
}
