package cli

import (
    "os"
    "errors"
    "github.com/thosebeans/wento/globals"
    "path"
)

func cliCommit() error {
    if len(os.Args) < 3 {
        return errors.New("not enough arguments for commit")
    } else if e := packageExists(os.Args[2]); e != nil {
        return e
    } else if e := os.Chdir(path.Join(globals.GetRootDir(), os.Args[2])); e != nil {
        return e
    } else if e := commandExec("git", "add", "-A"); e != nil {
        return e
    } else if len(os.Args) == 3 {
        return commandExec("git", "commit", "-m", "automated wento-commit")
    } else {
        return commandExec("git", "commit", "-m", os.Args[3])
    }
}
