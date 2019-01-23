package cli

import (
    "errors"
    "os/exec"
    "fmt"
)

func commandExec(c ...string) error {
    var cmd *exec.Cmd = exec.Command(c[0], c[1:]...)
    if s,e := cmd.CombinedOutput(); e != nil {
        return errors.New(fmt.Sprintf("%s\n%s", e.Error(), string(s)))
    } else {
        return nil
    }
}
