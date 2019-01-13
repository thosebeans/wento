package primitives

import (
    "os"
    "errors"
)

type srcPrimitive struct {
    src string
}

func (y srcPrimitive) testSrc() error {
    _,e := os.Open(y.src)
    return e
}

type dstPrimitive struct {
    dst string
}

func (y dstPrimitive) testDst() error {
    if y.dst == "" {
        return errors.New("no dst-path found")
    }
    return nil
}
