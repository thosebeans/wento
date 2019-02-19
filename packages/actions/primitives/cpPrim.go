package primitives

import (
    fs "github.com/thosebeans/wento/filesystem"
    "github.com/thosebeans/wento/globals"
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

func (y RawPrimitive) parseCp() (Primitive, error) {
    if len(y) < 2 {
        return nil,errors.New("cp: src missing")
    } else if len(y) < 3 {
        return nil,errors.New("cp: dst missing")
    }
    var p cpPrimitive = cpPrimitive{}
    if s,e0 := y[1].Template(globals.GetEnvs()); e0 != nil {
        return nil,e0
    } else if d,e1 := y[2].Template(globals.GetEnvs()); e1 != nil {
        return nil,e1
    } else {
        p.src = s
        p.dst = d
    }
    return p,nil
}
