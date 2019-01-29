package main

import (
    "os"
    "fmt"
)

func exitOne(formatStr string, Strs ...interface{}) {
    os.Stderr.Write([]byte(fmt.Sprintf(formatStr, Strs...)))
    os.Exit(1)
}
