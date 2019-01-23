package cli

import (
    "github.com/thosebeans/wento/globals"
)

func cliLs() error {
    if s,e := globals.GetPackages(); e != nil {
        return e
    } else {
        println("Packages:")
        for _,i := range s {
            println(i)
        }
        return nil
    }
}
