package cli

import (
    "os"
    "os/exec"
    "errors"
    "path"
    "github.com/thosebeans/wento/globals"
    "github.com/thosebeans/wento/filesystem"
)

func cliInit() error {
    if len(os.Args) < 3 {
        return errors.New("not enough arguments for init")
    } else if !nameCheck(os.Args[2]) {
        return errors.New("invalid package-name")
    } else if _,e := os.Open(path.Join(globals.GetRootDir(), os.Args[2])); e == nil {
        return errors.New("package already exists")
    } else if e := filesystem.MkdirP(path.Join(globals.GetRootDir(), os.Args[2])); e != nil {
        return e
    } else if e := os.Chdir(path.Join(globals.GetRootDir(), os.Args[2])); e != nil {
        return e
    }
    var cmd *exec.Cmd = exec.Command("git", "init")
    return cmd.Run()
}
