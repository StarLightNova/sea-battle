package playermap

import "fmt"

func mapPrinter(playerMap map[string][]string) string {
    returnString := fmt.Sprintf("     1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10\n")

    returnString += "===========================================\n"

    for _, key := range GetLetterCoordinates() {
        returnString += fmt.Sprintf(" %s )", key)

        for _, letter := range playerMap[key] {
            returnString += fmt.Sprintf(" %s |", letter)
        }

        returnString += "\n-------------------------------------------\n"
    }

    return returnString
}
