package main

import (
    "os/exec"
    "fmt"
    "errors"
)

func depCheck() error {
    var gitCmd *exec.Cmd = exec.Command("git", "--help")
    if e := gitCmd.Run(); e != nil {
        return errors.New(fmt.Sprintf("Dependecie-Error:\n%s", e.Error()))
    }
    var cpCmd *exec.Cmd = exec.Command("cp", "--help")
    if e := cpCmd.Run(); e != nil {
        return errors.New(fmt.Sprintf("Dependencie-Error:\n%s", e.Error()))
    }
    return nil
}
