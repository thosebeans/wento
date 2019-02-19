package primitives

import (
    fs "github.com/thosebeans/wento/filesystem"
    "errors"
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
