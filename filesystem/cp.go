package filesystem

import (
    "os"
    "os/exec"
    "errors"
    "fmt"
)

func CopyRF(src, dst string) error {
    if _,e := os.Open(src); e != nil {
        return e
    }
    if dst == "" {
        return errors.New("no dst-path given")
    }
    if e := RemoveRF(dst); e != nil {
        return e
    }
    var cmd *exec.Cmd = exec.Command("cp", "-rf", src, dst)
    if s,e := cmd.CombinedOutput(); e != nil {
        return errors.New(fmt.Sprintf(
            "%s\n%s", e.Error(), string(s),
        ))
    } else {
        return nil
    }
}
