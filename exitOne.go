package main

import (
    "os"
    "fmt"
)

func exitOne(formatStr string, Strs ...interface{}) {
    fmt.Printf(formatStr, Strs...)
    os.Exit(1)
}
