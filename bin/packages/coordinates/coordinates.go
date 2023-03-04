package coordinates

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Coordinates struct {
  StartRow string;
  StartColumn int;
  EndRow string;
  EndColumn int;
}

func New(strCordinates string) (*Coordinates, error) {
  regexPattern, _ := regexp.Compile("(\\w{1}\\d{1})(:| |-)(\\w{1}\\d{1})")

  if regexPattern.MatchString(strCordinates) {
    matchedGroupes := regexPattern.FindStringSubmatch(strCordinates)
    startGroup := strings.ToUpper(matchedGroupes[1])
    endGroup := strings.ToUpper(matchedGroupes[3])
    startColumn, _ := strconv.Atoi(startGroup[1:])
    endColumn, _ := strconv.Atoi(endGroup[1:])

    return &Coordinates{
      StartRow: startGroup[:1],
      StartColumn: startColumn,
      EndRow: endGroup[:1],
      EndColumn: endColumn,
    }, nil
  }

  return nil, errors.New("The given coordinates are incorrect.")
}

func (coor Coordinates) String() string {
  return fmt.Sprintf("%s%d:%s%d", coor.StartRow, coor.StartColumn, coor.EndRow, coor.EndColumn)
}
