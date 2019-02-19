package primitives

import (
    "os"
    "os/exec"
    "path"
    "fmt"
    "errors"
    "github.com/thosebeans/wento/globals"
)

func init() {
    primitives["run"] = parseRunPrim
}

type runPrimitive struct {
    srcPrimitive
}

func (y runPrimitive) Test() error {
    return y.testSrc()
}

func (y runPrimitive) Emerge() error {
    var cmd *exec.Cmd
    if path.IsAbs(y.src) {
        cmd = exec.Command(y.src)
    } else {
        cmd = exec.Command(fmt.Sprintf(
            "%s%s", "./", y.src,
        ))
    }
    cmd.Stdout = os.Stdout
    cmd.Stdin  = os.Stdin
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func parseRunPrim(s []string) (Primitive, error) {
    if len(s) == 1 {
        return nil,errors.New("src or cmd missing")
    }
    var p runPrimitive
    p.src = s[1]
    return p,nil
}

func (y RawPrimitive) parseRun() (Primitive, error) {
    if len(y) < 2 {
        return nil,errors.New("run: src missing")
    }
    var p runPrimitive = runPrimitive{}
    if s,e := y[1].Template(globals.GetEnvs()); e != nil {
        return nil,e
    } else {
        p.src = s
    }
    return p,nil
}
