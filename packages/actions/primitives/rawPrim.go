package primitives

import (
    "github/thosebeans/wento/packages/sets/actions/primitives/templateString"
    "errors"
    "fmt"
)

type RawPrimitive []templateString.TemplateString

func (y RawPrimitive) Parse() (Primitive, error) {
    if len(y) == 0 {
        return nil,errors.New("primitive contains no data")
    }
    switch y[0] {
    case "ln":
        return y.parseLn()
    case "cp":
        return y.parseCp()
    case "run":
        return y.parseRun()
    default:
        return nil,errors.New(fmt.Sprintf(
            "%s %s", y[0], "isnt a proper primitive",
        ))
    }
}
