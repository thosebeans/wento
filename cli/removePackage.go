package cli

import (
    "os"
    "errors"
    "path"
    "github.com/thosebeans/wento/globals"
)

func cliRemove() error {
    if len(os.Args) < 3 {
        return errors.New("not enough arguments for remove")
    } else if _,e := os.Open(path.Join(globals.GetRootDir(), os.Args[2])); e != nil {
        return e
    } else {
        return os.RemoveAll(path.Join(globals.GetRootDir(), os.Args[2]))
    }
}
