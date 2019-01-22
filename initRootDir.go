package main

import (
    fs "github.com/thosebeans/wento/filesystem"
    "github.com/thosebeans/wento/globals"
)

func initRootDir() error {
    return fs.MkdirP(globals.GetRootDir())
}
