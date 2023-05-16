package playermap

const (
    Width int = 10
    Height    = Width
)

const (
    EMPTY string = " "
    HIT          = "X"
    SHIP         = "S"
    MISS         = "M"
)

const (
    MAP_COLUMNS string    = "     1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10"
    COLUMN_BOTTOM_DIVIDER = "==========================================="
    BOTTOM_DIVIDER        = "-------------------------------------------"
    YOUR_BOARD_NAME       = "                   YOU                     "
    OPPONENT_BOARD_NAME   = "                 OPPONENT                  "
)

func GetLetterCoordinates() [Width]string {
    return [Width]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
}

func emptyRow() []string {
    return []string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "}
}

