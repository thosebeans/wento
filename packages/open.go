package packages

import (
    "github.com/thosebeans/wento/globals"
    "path"
    "os"
    "io/ioutil"
)

func OpenPackage(name string) (Package, error) {
    var filePath string = path.Join(globals.GetRootDir(), name, "package.json")
    if f,e := ioutil.ReadFile(filePath); e != nil {
        return Package{},e
    } else {
        defer os.Chdir(path.Join(globals.GetRootDir(), name))
        return parsePackage(f)
    }
}
