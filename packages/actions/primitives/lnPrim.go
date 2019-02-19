package primitives

import (
    fs "github.com/thosebeans/wento/filesystem"
    "errors"
    "github.com/thosebeans/wento/globals"
    "path"
)

func init() {
    primitives["ln"] = parseLnPrim
}

type lnPrimitive struct {
    srcPrimitive
    dstPrimitive
}

func (y lnPrimitive) Test() error {
    if e := y.testSrc(); e != nil {
        return e
    }
    return y.testDst()
}

func (y lnPrimitive) Emerge() error {
    if e := fs.MkdirP(path.Dir(y.dst)); e != nil {
        return e
    } else {
        return fs.LinkSF(y.src, y.dst)
    }
}

func parseLnPrim(s []string) (Primitive, error) {
    switch len(s) {
    case 1:
        return nil,errors.New("ln: src missing")
    case 2:
        return nil,errors.New("ln: dst missing")
    }
    var p lnPrimitive
    p.src = s[1]
    p.dst = s[2]
    return p,nil
}

func (y RawPrimitive) parseLn() (Primitive, error) {
    if len(y) < 2 {
        return nil,errors.New("ln: src missing")
    } else if len(y) < 3 {
        return nil,errors.New("ln: dst missing")
    }
    var p lnPrimitive = lnPrimitive{}
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
