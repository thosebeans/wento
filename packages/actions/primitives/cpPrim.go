package primitives

import (
    fs "github.com/thosebeans/wento/filesystem"
    "errors"
    "path"
)

func init() {
    primitives["cp"] = parseCpPrim
}

type cpPrimitive struct {
    lnPrimitive
}

func (y cpPrimitive) Emerge() error {
    if e := fs.MkdirP(path.Dir(y.dst)); e != nil {
        return e
    } else {
        return fs.CopyRF(y.src, y.dst)
    }
}

func parseCpPrim(s []string) (Primitive, error) {
    switch len(s) {
    case 2:
        return nil,errors.New("cp: src missing")
    case 3:
        return nil,errors.New("cp: dst missing")
    }
    var p cpPrimitive
    p.src = s[1]
    p.dst = s[2]
    return p,nil
}
