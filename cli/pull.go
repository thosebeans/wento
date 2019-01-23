package cli

import (
    "os"
    "errors"
    "github.com/thosebeans/wento/globals"
    "path"
)

func cliPull() error {
    if len(os.Args) < 3 {
        return errors.New("not enough arguments for pull")
    } else if e := packageExists(os.Args[2]); e != nil {
        return e
    } else if e := os.Chdir(path.Join(globals.GetRootDir(), os.Args[2])); e != nil {
        return e
    }
    return commandExec("git", "pull")
}
