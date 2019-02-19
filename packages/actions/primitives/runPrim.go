package primitives

import (
    "os"
    "os/exec"
    "errors"
)

func init() {
    primitives["run"] = parseRunPrim
}

type runPrimitive struct {
    exe  string
    args []string
}

func (y runPrimitive) Test() error {
    if _,e := exec.LookPath(y.exe); e == nil  {
        return nil
    } else if _,e := os.Open(y.exe); e == nil {
        return nil
    } else {
        return errors.New("no proper executable found")
    }
}

func (y runPrimitive) Emerge() error {
    var c *exec.Cmd = exec.Command(y.exe, y.args...)
    c.Stdout = os.Stdout
    c.Stdin  = os.Stdin
    c.Stderr = os.Stderr
    return c.Run()
}

func parseRunPrim(s []string) (Primitive, error) {
    if len(s) == 1 {
        return nil,errors.New("src or cmd missing")
    }
    var p runPrimitive
    p.exe = s[1]
    if len(s) > 2 {
        p.args = s[2:]
    }
    return p,nil
}
