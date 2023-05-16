package playermap

type PlayerMap struct {
    TheMap map[string][]string
    ShadowMap map[string][]string
}

func New() (*PlayerMap, error) {
    pureMap := make(map[string][]string)
    shadowMap := make(map[string][]string)

    fillPureMap(pureMap)
    fillPureMap(shadowMap)

    return &PlayerMap {
        TheMap: pureMap,
        ShadowMap: shadowMap,
    }, nil
}

func (playerMap PlayerMap) String() string {
    return mapPrinter(playerMap.TheMap)
}

func (playerMap PlayerMap) ShadowString() string {
    return mapPrinter(playerMap.ShadowMap)
}

