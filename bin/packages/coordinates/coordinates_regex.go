package coordinates

import (
    "errors"
    "regexp"
    "strconv"
    "strings"
)

const (
    CoordinatesPattern string = "(\\w{1}(10|\\d{1}))(:| |-)(\\w{1}(10|\\d{1}))"
    HalfCoordinatesPattern    = "(\\w{1})(10|\\d{1})"
)

func StringToCoordinates(strCoordinates string) (Coordinates, error) {
    regexPattern, _ := regexp.Compile(CoordinatesPattern)

    if regexPattern.MatchString(strCoordinates) {
        matchedGroupes := regexPattern.FindStringSubmatch(strCoordinates)

        startGroup := strings.ToUpper(matchedGroupes[1])
        endGroup := strings.ToUpper(matchedGroupes[4])
        startColumn, _ := strconv.Atoi(startGroup[1:])
        endColumn, _ := strconv.Atoi(endGroup[1:])

        return Coordinates{
            StartRow: startGroup[:1],
            StartColumn: startColumn,
            EndRow: endGroup[:1],
            EndColumn: endColumn,
        }, nil
    }

    return Coordinates{}, errors.New("The given coordinates are incorrect.")
}

func StringToRowColumn(userCoordindates string) (row string, column int) {
    regexPattern, _ := regexp.Compile(HalfCoordinatesPattern)

    if regexPattern.MatchString(userCoordindates) {
        matchedGroupes := regexPattern.FindStringSubmatch(userCoordindates)

        row = strings.ToUpper(matchedGroupes[1])
        column, _ = strconv.Atoi(matchedGroupes[2])
    } else {
        panic("The half coordinates (user's input) is not correct or formatter error.")
    }

    return
}

