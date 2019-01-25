# wento
wento is a package-manager and automatation-framework, for your configurations.  
--
If you want to...
- install your dotfiles on a new machine
- share your configuration with others
- easily sync your dotfiles to a git-repo

**wento** got you covered.


## Installation
Dependencies:
- **Go** (the program is developed with Go 1.11.4), only for compilation from source
- a posix conformant version of **cp**
- **git**

Install-Process:
- execute `go install github.com/thosebeans/wento` in your shell
- if your `$GOBIN` is contained in your path, you just installed it, else, copy `$GOPATH/bin/wento` into a directory of your path

## Usage
### Packages
- packages are stored in `~/.wento`
- packages are git-repositories
- packages contain a `package.json`-file, that defines all information, that is used by wento
### Actions
- packages contain an undefined amount of actions
- actions have a description and and undefined amount of primitives
### Primitives
- primitives are lower-level operations
- primtives get executed, when en action is emerged

## package.json
```json
{
  "desc": "mydotfiles",
  "actions": {
    "bash": {
      "desc": "bash-dotfiles",
      "primitives": [
        ["ln", "bash/bashrc", "{{.HOME}}/.bashrc"],
        ["ln", "bash/profile", "{{.HOME}}/.profile"]
      ]
    },
    "i3": {
      "desc": "i3-files",
      "primitives": [
        ["ln", "i3/i3con", "{{.HOME}}/.config/i3/config"],
        ["ln", "i3/i3stat", "{{.HOME}}/.config/i3status/config"]
      ]
    }
  }
}
```

**Note:** `desc`-attributes are currently not used

### string-templating
wento uses Go's default templating-engine to include environment-variables in your paths  
You can use them like this:  
- `{{.HOME}}/.bashrc` evaluates to `/home/username/.bashrc`  
- `{{.XDG_CONFIG_HOME}}/nvim` evaluates to `/home/username/.config/nvim`  
**Note:** Environment-Variables are fetched and cached, the first time, when wento needs to use one.

### primitives
- primitives are string-arrays
- relative paths are matched relative to the package-directory
- all file-system-primitives are **FORCE**, so they will override the target, if it exists
- if a primitives requires a file, the existence of this file is tested, before it gets emerged

| Primitive | Example | Description |
|--|--|--|
| Link | `["ln", "fileA", "fileB"` | Creates a symlink `fileB` that points to `fileA` |
| Copy | `["cp", "fileA", "fileB"]`| Copy `fileA` to `fileB` |
| Run | `["run", "fileA"]` | Execute `fileA` |

**NOTE:** All primitives are executed from inside the package-directory

## cli
| Command | Description |
|--|--|
| `wento help` | Shows a help-message |
| `wento clone <URL>` | Clones the remote-repository from `<URL>` into `~/.wento` |
| `wento commit <PKG> <MSG>` | Commits all changes to `<PKG>`. If `<MSG>` is omitted, a default commit-message is used |
| `wento push <PKG>` | Push changes of `<PKG>` to its default remote-repository |
| `wento pull <PKG>` | Pull changes of `<PKG>` from its default remote-repository |
| `wento init <PKG>` | Inits a new git-repository called `<PKG>` in `~/.wento` |
| `wento remove <PKG>` | Removes `<PKG>` from `~/.wento` |
| `wento info <PKG> <ACT>` | Show information of <ACT> in <PKG>, if <ACT> is omitted, information of <PKG> is shown |
| `wento list` | Lists all packages |
| `wento emerge <PKG> <ACT>` | Emerges all primitives of `<ACT>` in `<PKG>` |
