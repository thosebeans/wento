package cli

import (
    "os"
    "errors"
    "github/thosebeans/wento/globals"
    "os/exec"
    "path"
)

func cliPush() error {
    if len(os.Args) < 3 {
        return errors.New("not enough arguments for push")
    } else if e := os.Chdir(path.Join(globals.GetRootDir(), os.Args[2])); e != nil {
        return e
    }
    var cmd *exec.Cmd = exec.Command("git", "push")
    return cmd.Run()
}