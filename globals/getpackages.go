package globals

import (
    "os"
)

func GetPackages() ([]string,error) {
    if f,e := os.Open(GetRootDir()); e != nil {
        return []string{},e
    } else {
        return f.Readdirnames(-1)
    }
}
