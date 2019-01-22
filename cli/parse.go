package cli

import (
    "os"
    "errors"
)

func ParseCli() error {
    if len(os.Args) < 2 {
        return errors.New("no arguments given")
    }
    var e error
    switch os.Args[1] {
    case "emerge":
        return cliEmerge()
    case "list":
        return cliLs()
    case "remove":
        return cliRemove()
    case "init":
        return cliInit()
    case "help":
        return cliHelp()
    case "clone":
        return cliClone()
    case "commit":
        return cliCommit()
    default:
        return errors.New("no proper arguments given")
    }
    return e
}
