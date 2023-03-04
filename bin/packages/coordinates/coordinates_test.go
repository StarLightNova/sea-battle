package coordinates

import (
  "testing"
)

func TestNewCoordinates(t *testing.T) {
  newCoor, _ := New("A1:A3")

  if newCoor.startRow != "A" || newCoor.endColumn != 3 || newCoor.endRow != "A" || newCoor.startColumn != 1 {
    t.Fatal("The coordinate constructor gave incorrect coordinates")
  }
}

func TestIncorrectNewCoordinates(t *testing.T) {
  failCounter := 0
  failCoordinates := []string {
    "bi:bi",
    "1B 8*",
    "mm---kek",
  }

  for _, failCoordinate := range failCoordinates {
    _, error := New(failCoordinate)

    if error != nil {
      failCounter++
    }
  }

  if failCounter != len(failCoordinates) {
    t.Fatal("The coordinate constructor is failing.")
  }
}
