package cli

const helpStr string = `
wento help          - show this
wento init   <NAME> - create a new git-repo in ~/.wento/<NAME>
wento remove <NAME> - remove package <NAME>
wento clone  <URL>  - clone a git-repo in ~/.wento
wento commit <PKG> <MSG> - commits all changes for <PKG>, <MSG> may be omitted
wento pull   <PKG>  - pull changes for <PKG>
wento push   <PKG>  - push changes of <PKG> to its remote repository
wento list          - list local packages
wento emerge <PKG> <ACTION> - emerge <ACTION> from <PKG>
`

func cliHelp() error {
    println(helpStr)
    return nil
}
