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
    default:
        return errors.New("no proper arguments given")
    }
    return e
}
