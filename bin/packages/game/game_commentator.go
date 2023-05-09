package game

import (
	"fmt"
    "time"
	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
)

func announceAGame() {
    time.Sleep(3 * time.Second)
    fmt.Println("\nWelcome to the Sea Battle!\nI hope both of you can swim.")
    time.Sleep(2 * time.Second)
    fmt.Println("\nRules can be accesses by command in terminal prompt: [cat rules.txt] (without '[' and ']')")
    time.Sleep(2 * time.Second)
}

func yourBoard(playermap playermap.PlayerMap) {
    fmt.Println("Generating a random board with ships...")
    time.Sleep(1 * time.Second)
    fmt.Println("\nYour randomly generated board looks like this:")
    fmt.Println(playermap)
}

func opponentsShadowMap(playermap playermap.PlayerMap) {
    fmt.Println("\nYou are attacking OPPONENTS board:")
    fmt.Println(playermap.ShadowString())
}

func yourBoardHasAttacked(playermap playermap.PlayerMap) {
    fmt.Println("\n!!!Your board has attacked!!!")
    fmt.Println(playermap)
}

func howToAttackCoordinates() {
    fmt.Println("\n [INFO] You can attack prompting like these examples: A1, B6 etc,.")
}

func readyToGo() {
    fmt.Println("Ready? (y/n)")
}

func goodbye() {
    fmt.Println("Bye bye 👋")
}

func clearConsole() {
    fmt.Println("\x1bc")
}
