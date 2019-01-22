package packages

import (
    "github.com/thosebeans/wento/packages/actions"
    "encoding/json"
    "errors"
    "fmt"
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

func (y Package) GetAction(name string) (actions.Action, error) {
    if y.Actions == nil || len(y.Actions) == 0 {
        return actions.Action{},errors.New("package doesnt contain any actions")
    } else if a := y.Actions[name]; a.Desc == "" && len(a.Primitives) == 0 {
        return a,errors.New(fmt.Sprintf("%s %s", "package doesnt contain an action:", name))
    } else {
        return a,nil
    }
}
