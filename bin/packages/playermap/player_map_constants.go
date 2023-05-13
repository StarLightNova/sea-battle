package playermap

const (
    Width int = 10
    Height = Width
)

func GetLetterCoordinates() [Width]string {
    return [Width]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
}

func emptyRow() []string {
    return []string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " "}
}

