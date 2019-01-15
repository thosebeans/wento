package actions

import (
    "github/thosebeans/wento/packages/sets/actions/primitives"
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

func (y Action) Test() ([]string,error) {
    var prims []primitives.Primitive
    var errs  []string
    if p,e := y.parse(); e != nil {
        return errs,e
    } else {
        prims = p
    }
    for _,i := range prims {
        if e := i.Test(); e != nil {
            errs = append(errs, e.Error())
        }
    }
    return errs,nil
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
