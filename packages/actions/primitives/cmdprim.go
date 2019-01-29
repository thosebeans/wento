package primitives

import (
    "errors"
    "os/exec"
    "os"
    "github.com/thosebeans/wento/globals"
)

type cmdPrim []string

func (y RawPrimitive) parseCmd() (Primitive, error) {
    if len(y) < 2 {
        return nil,errors.New("not enough arguments for cmd")
    }
    var p cmdPrim
    for _,v := range y[1:] {
        if s,e := v.Template(globals.GetEnvs()); e != nil {
            return p,e
        } else {
            p = append(p, s)
        }
    }
    return p,nil
}

func (y cmdPrim) Test() error {
    return nil
}

func (y cmdPrim) Emerge() error {
    var cmd *exec.Cmd
    if len(y) == 1 {
        cmd = exec.Command(y[0])
    } else {
        cmd = exec.Command(y[0], y[1:]...)
    }
    cmd.Stdin  = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}
