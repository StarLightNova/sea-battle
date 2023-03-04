package ship

import (
  "testing"
  // "fmt"
)

func TestObjectShip(t *testing.T) {
  _, error := New("Carrier")

  if error != nil {
    t.Fatal("Can not create a new Ship. Constructor error.")
  }
}

func TestNewShipNoneClass(t *testing.T) {
  newShip, error := New("carrier")

  if newShip != nil &&  error == nil  {
    t.Fatal("The constructor of the Ship returned actual ship, instead of the erorr.")
  }
}
