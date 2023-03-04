package coordinates

import (
  "regexp"
  "strconv"
  "errors"
  "fmt"
)

type Coordinates struct {
  startRow string;
  startColumn int;
  endRow string;
  endColumn int;
}

func New(strCordinates string) (Coordinates, error) {
  regexPattern, _ := regexp.Compile("(\\w{1}\\d{1})(:| |-)(\\w{1}\\d{1})")

  if regexPattern.MatchString(strCordinates) {
    matchedGroupes := regexPattern.FindStringSubmatch(strCordinates)
    startGroup := matchedGroupes[1]
    endGroup := matchedGroupes[3]
    startColumn, _ := strconv.Atoi(startGroup[1:])
    endColumn, _ := strconv.Atoi(endGroup[1:])

    return Coordinates{
      startRow: startGroup[:1],
      startColumn: startColumn,
      endRow: endGroup[:1],
      endColumn: endColumn,
    }, nil
  }

  return Coordinates{}, errors.New("The given coordinates are incorrect.")
}

func (coor Coordinates) String() string {
  return fmt.Sprintf("%s%d:%s%d", coor.startRow, coor.startColumn, coor.endRow, coor.endColumn)
}
