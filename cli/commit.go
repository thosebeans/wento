package cli

import (
    "os"
    "errors"
    "github.com/thosebeans/wento/globals"
    "os/exec"
    "path"
)

func cliCommit() error {
    if len(os.Args) < 3 {
        return errors.New("not enough arguments for commit")
    } else if e := os.Chdir(path.Join(globals.GetRootDir(), os.Args[2])); e != nil {
        return e
    }
    var cmd *exec.Cmd = exec.Command("git", "add", "-A")
    if e := cmd.Run(); e != nil {
        return e
    }
    if len(os.Args) == 3 {
        cmd = exec.Command("git", "commit", "-m", "automated wento-commit")
    } else {
        cmd = exec.Command("git", "commit", "-m", os.Args[3])
    }
    return cmd.Run()
}
