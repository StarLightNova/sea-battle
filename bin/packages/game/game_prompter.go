package game

import (
    "fmt"
    "strings"
)

func promptToStartAGame() bool {
    var usersInput string

    for {
        fmt.Scanln(&usersInput)
        loweredUsersPrompt := strings.ToLower(usersInput)

        if loweredUsersPrompt == "y" {
            return true
        } else if loweredUsersPrompt == "n" {
            return false
        } else {
            fmt.Println("Your reply is not understandable by computer. (y/n)")
        }
    }
}

