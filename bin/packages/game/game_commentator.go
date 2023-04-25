package game

import (
	"fmt"
	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
)

func announceAGame() {
    fmt.Println("\nWelcome to the Sea Battle!\nI hope both of you can swim.")
    fmt.Println("\nRules can be accesses by command: [cat rules.txt] (without '[' and ']')")
}

func yourBoard(playermap playermap.PlayerMap) {
    fmt.Println("\nYour randomly generated board look like this:")
    fmt.Println(playermap)
}

func readyToGo() {
    fmt.Println("Ready? (y/n)")
}

func goodbye() {
    fmt.Println("Bye bye 👋")
}
