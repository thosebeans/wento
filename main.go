package main

import (
    "github.com/thosebeans/wento/cli"
)

func main() {
    if e := initRootDir(); e != nil {
        exitOne("%s\n", e.Error())
    }
    if e := cli.ParseCli(); e != nil {
        exitOne("%s\n", e.Error())
    }
}
