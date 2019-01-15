package packages

import (
    "github/thosebeans/wento/packages/actions"
    "encoding/json"
)

type Package struct {
    Desc    string                    `json:"desc"`
    Actions map[string]actions.Action `json:"actions"`
}

func parsePackage(jsonBytes []byte) (Package, error) {
    var p Package
    var e error = json.Unmarshal(jsonBytes, &p)
    return p,e
}
