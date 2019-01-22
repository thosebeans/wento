package cli

import (
    "os"
    "errors"
    "github.com/thosebeans/wento/globals"
    "os/exec"
    "path"
)

func cliPull() error {
    if len(os.Args) < 3 {
        return errors.New("not enough arguments for pull")
    } else if e := os.Chdir(path.Join(globals.GetRootDir(), os.Args[2])); e != nil {
        return e
    }
    var cmd *exec.Cmd = exec.Command("git", "pull")
    return cmd.Run()
}
