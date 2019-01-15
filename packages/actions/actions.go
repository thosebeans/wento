package actions

import (
    "github/thosebeans/wento/packages/actions/primitives"
    "errors"
    "fmt"
)

type Action struct {
    Desc string                          `json:"desc"`
    Primitives []primitives.RawPrimitive `json:"primitives"`
}

func (y Action) parse() ([]primitives.Primitive, error) {
    var prims []primitives.Primitive
    for _,r := range y.Primitives {
        if p,e := r.Parse(); e != nil {
            return prims,e
        } else {
            prims = append(prims, p)
        }
    }
    return prims,nil
}

func (y Action) Test() error {
    var prims []primitives.Primitive
    var errs  []string
    if p,e := y.parse(); e != nil {
        return e
    } else {
        prims = p
    }
    for _,i := range prims {
        if e := i.Test(); e != nil {
            errs = append(errs, e.Error())
        }
    }
    if len(errs) > 0 {
        var errStr string
        for _,i := range errs {
            errStr = fmt.Sprintf("%s\n%s", errStr, i)
        }
        return errors.New(errStr)
    } else {
        return nil
    }
}

func (y Action) Emerge() error {
    var prims []primitives.Primitive
    if p,e := y.parse(); e != nil {
        return e
    } else {
        prims = p
    }
    for _,i := range prims {
        if e := i.Emerge(); e != nil {
            return e
        }
    }
    return nil
}
