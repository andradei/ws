# ws

CLI workspace manager. Accessible via the `ws` binary

This tool outputs the path of a given workspace so your shell can use to change directories. This is
necessary because processes change working directory only in the context in which they run, the
directory changes back to what it was before the program ran when it exits.

## Installation

1. Run on your terminal `go get github.com/andradei/ws` (Homebrew formula will be available later)
1. Integrate `ws` in your shell with the following function on your `.bashrc` file or equivalent:
    ```bash
    # A function with the same name as the program.
    function ws {
        # Replace the value below by the path to the ws binary on your machine.
        local path=$GOPATH/bin/ws
        # Call ws and pass to it all the arguments given to this function by using $@.
        if dir=$($path $@); then
            # If successful, ws will output the directory of a workspace. So cd into that.
            cd "$dir"
        fi
    }
    ```
1. Reload your shell (easiest way is restarting your terminal) `source ~/.bashrc` (or equivalent to your shell configuration file).

## Commands

TODO: Create output of `ws -help` and add it below

TODO: Create a gif with usage examples

```bash
ws [command | workspace name] [workspace name]

-insert | -i <name>    Create a workspace with name
-delete | -d <name>    Delete an existing workspace with name
-help | -h             Display this help message
-list | -l             List existing workspaces
```

## Examples

- Create workspace: `ws -insert project1`

- Go to workspace: `ws project1`

- Delete workspace: `ws -delete project1`

## Behavior

- Can't override existing workspaces (`-force` option will be added later to change this on demand)
- Multiple workspaces can have the same path (aliasing)
- TODO: Workspace named _default_ can be accessed without parameters, just run `ws`

## Implementation Details

- `fmt.Print` is used when the output of `ws` is a directory path. Example: when running `ws my-workspace`
- `printMsg` is used when the output is an error or some information. Example: when running `ws -list`