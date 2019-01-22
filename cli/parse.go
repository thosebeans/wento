package cli

import (
    "os"
    "errors"
)

func ParseCli() error {
    if len(os.Args) < 2 {
        return errors.New("no arguments given")
    }
    return nil
}
