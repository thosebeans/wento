package actions

import (
    "github/thosebeans/wento/packages/sets/actions/primitives"
)

type Action struct {
    Desc string                          `json:"desc"`
    Primitives []primitives.RawPrimitive `json:"primitives"`
}
}
