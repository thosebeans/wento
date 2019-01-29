package primitives

import (
    "os"
    "os/exec"
    "errors"
    "github.com/thosebeans/wento/globals"
)

type shellPrimitivePosix struct {
    shell string
    input string
}

func (y shellPrimitivePosix) Test() error {
    _,e := exec.LookPath(y.shell)
    return e
}

func (y shellPrimitivePosix) Emerge() error {
    var cmd *exec.Cmd = exec.Command(y.shell, "-c", y.input)
    cmd.Stdin  = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func (y RawPrimitive) parseShellPosix() (Primitive, error) {
    if shell,e := y[1].Template(globals.GetEnvs()); e != nil {
        return nil,e
    } else if input,e := y[2].Template(globals.GetEnvs()); e != nil {
        return nil,e
    } else {
        var p shellPrimitivePosix = shellPrimitivePosix{}
        p.input = input
        p.shell = shell
        return p,nil
    }
}

func (y RawPrimitive) parseShell() (Primitive, error) {
    if len(y) < 3 {
        return nil,errors.New("not enough arguments for shell")
    } else if y[1] == "bash" || y[1] == "sh" || y[1] == "ksh" || y[1] == "zsh" {
        return y.parseShellPosix()
    }
    switch y[1] {
    default:
        return nil,errors.New("no proper shell found")
    }
}
