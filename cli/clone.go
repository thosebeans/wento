package cli

import (
    "os"
    "errors"
    "github.com/thosebeans/wento/globals"
    "os/exec"
)

func cliClone() error {
    if len(os.Args) < 3 {
        return errors.New("not enough arguments for clone")
    } else if e := os.Chdir(globals.GetRootDir()); e != nil {
        return e
    }
    var cmd *exec.Cmd = exec.Command("git", "clone", os.Args[2])
    return cmd.Run()
}
