package main

import (
    fs "github/thosebeans/wento/filesystem"
    "github/thosebeans/wento/globals"
)

func initRootDir() error {
    return fs.MkdirP(globals.GetRootDir())
}
