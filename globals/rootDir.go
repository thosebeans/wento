package globals

import (
    "path"
    "os"
)

func GetRootDir() string {
    return path.Join(
        os.Getenv("HOME"), ".wento",
    )
}
