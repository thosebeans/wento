package filesystem

import (
    "os"
    "errors"
    "path"
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
    if path.IsAbs(src) {
        return os.Symlink(src, dst)
    } else if s,e := os.Getwd(); e != nil {
        return e
    } else {
        return os.Symlink(path.Join(s, src), dst)
    }
}
