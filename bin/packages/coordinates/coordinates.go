package coordinates

import "fmt"

type Coordinates struct {
    StartRow string;
    StartColumn int;
    EndRow string;
    EndColumn int;
}

func New(strCoordinates string) (*Coordinates, error) {
    coordinates, error := StringToCoordinates(strCoordinates)

    if error != nil {
        return nil, error
    }

    return &coordinates, nil
}

func (coor Coordinates) String() string {
    return fmt.Sprintf("%s%d:%s%d", coor.StartRow, coor.StartColumn, coor.EndRow, coor.EndColumn)
}
