package primitives

import (
    "github.com/thosebeans/wento/packages/actions/primitives/templateString"
    "github.com/thosebeans/wento/globals"
    "errors"
)

type RawPrimitive []templateString.TemplateString

func (y RawPrimitive) Parse() (Primitive, error) {
    var primSlice []string
    if len(y) == 0 {
        return nil,errors.New("primitive contains no data")
    } else {
        primSlice = make([]string, len(y))
        for i,v := range y {
            if s,e := v.Template(globals.GetEnvs()); e != nil {
                return nil,e
            } else {
                primSlice[i] = s
            }
        }
    }
    if f := primitives[(primSlice[0])]; f != nil {
        return f(primSlice)
    }
    return nil,errors.New("no proper primitive found")
}
