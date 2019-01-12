package filesystem

import (
    "os"
)

func RemoveRF(path string) error {
    return os.RemoveAll(path)
}
