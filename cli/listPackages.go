package cli

import (
    "os"
    "github.com/thosebeans/wento/globals"
)

func cliLs() error {
    if d,e := os.Open(globals.GetRootDir()); e != nil {
        return e
    } else if c,e := d.Readdirnames(-1); e != nil {
        return e
    } else {
        println("Packages:")
        for _,i := range c {
            println(i)
        }
        return nil
    }
}
