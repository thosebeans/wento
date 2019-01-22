package cli

import (
    "os"
    "path"
    "errors"
    "github.com/thosebeans/wento/globals"
    "github.com/thosebeans/wento/packages"
)

func cliEmerge() error {
    if len(os.Args) < 4 {
        return errors.New("no enough arguments for emerge")
    } else if p,e := packages.OpenPackage(os.Args[2]); e != nil {
        return e
    } else if a,e := p.GetAction(os.Args[3]); e != nil {
        return e
    } else if e := os.Chdir(path.Join(globals.GetRootDir(), os.Args[2])); e != nil {
        return e
    } else if e := a.Test(); e != nil {
        return e
    } else {
        return a.Emerge()
    }
}
