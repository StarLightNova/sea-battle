package game

import (
	"fmt"
    "time"
	"github.com/StarLightNova/sea-battle/bin/packages/playermap"
)

func announceAGame() {
    fmt.Print("\nWelcome to the Sea Battle! ")
    time.Sleep(2 * time.Second)
    fmt.Println("I hope both of you can swim.")

    time.Sleep(2 * time.Second)
    fmt.Print("\nRules can be accesses by command in the terminal prompt: ")
    time.Sleep(2 * time.Second)
    fmt.Println("> cat rules.txt")
}

func yourBoard(playermap playermap.PlayerMap) {
    fmt.Println("Generating a random board with ships...")
    time.Sleep(1 * time.Second)
    fmt.Println("\nYour randomly generated board:")
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
    fmt.Println("\n [INFO] You can attack prompting like these examples: A1, B6, c3, j10 etc,.")
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
