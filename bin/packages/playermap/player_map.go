package playermap

type PlayerMap struct {
    theMap map[string][]string
    ShadowMap map[string][]string
}

func New() (*PlayerMap, error) {
    pureMap := make(map[string][]string)
    shadowMap := make(map[string][]string)

    fillPureMap(pureMap)
    fillPureMap(shadowMap)

    return &PlayerMap {
        theMap: pureMap,
        ShadowMap: shadowMap,
    }, nil
}

func (playerMap PlayerMap) String() string {
    return mapPrinter(playerMap.theMap)
}

func (playerMap PlayerMap) ShadowString() string {
    return mapPrinter(playerMap.ShadowMap)
}

