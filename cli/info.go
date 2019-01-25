package cli

import (
    "errors"
    "os"
    "github.com/thosebeans/wento/packages"
    "fmt"
)

func cliInfo() error {
    if len(os.Args) < 3 {
        return errors.New("not enough arguments for info")
    } else if e := packageExists(os.Args[2]); e != nil {
        return e
    } else if p,e := packages.OpenPackage(os.Args[2]); e != nil {
        return e
    } else if len(os.Args) < 4 {
        fmt.Printf("Description: %s\n\nActions:\n", p.Desc)
        for i := range p.Actions {
            println(i)
        }
        return nil
    } else if a,e := p.GetAction(os.Args[3]); e != nil {
        return e
    } else {
        fmt.Printf("Package: %s\nAction: %s\nDescription: %s\n\nPrimitives:\n", os.Args[2], os.Args[3], a.Desc)
        for _,i := range a.Primitives {
            fmt.Println(i)
        }
        return nil
    }
}
