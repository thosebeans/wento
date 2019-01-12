package filesystem

import (
    "os"
)

func MkdirP(path string) error {
    return os.MkdirAll(path, os.ModeDir | 0755)
}
