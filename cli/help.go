package cli

const helpStr string = `
wento help          - show this
wento init   <NAME> - create a new git-repo in ~/.wento/<NAME>
wento remove <NAME> - remove package <NAME>
wento list          - list local packages
wento emerge <PKG> <ACTION> - emerge <ACTION> from <PKG>
`

func cliHelp() error {
    println(helpStr)
    return nil
}
