package filesystem

import (
    "os"
    "errors"
)

func LinkSF(src, dst string) error {
    if _,e := os.Open(src); e != nil {
        return e
    }
    if dst == "" {
        return errors.New("no dst-path given")
    }
    if e := RemoveRF(dst); e != nil {
        return e
    }
    return os.Symlink(src, dst)
}
