package actions

import (
    "github/thosebeans/wento/packages/sets/actions/primitives"
)

type Action struct {
    desc string
    primitives []primitives.RawPrimitive
}
