package cli

import (
    "regexp"
)

func nameCheck(name string) bool {
    if name == "." || name == ".." {
        return false
    } else if m,e := regexp.MatchString(`^(\w|\d|[.-])+$`, name); !m || e != nil {
        return false
    } else {
        return true
    }
}
