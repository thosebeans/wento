package globals

import (
    "os"
    "strings"
)

var envVariables map[string]string

func GetEnvs() map[string]string {
    if envVariables != nil {
        return envVariables
    }
    envVariables = map[string]string{}
    for _,i := range os.Environ() {
        var s []string = strings.Split(i, "=")
        envVariables[(s[0])] = s[1]
    }
    return envVariables
}
