package cli

import (
    "errors"
    "github.com/thosebeans/wento/globals"
)

func packageExists(pkg string) error {
    if p,e := globals.GetPackages(); e != nil {
        return e
    } else {
        for _,i := range p {
            if i == pkg {
                return nil
            }
        }
        return errors.New("package not found")
    }
}
