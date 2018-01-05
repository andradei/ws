# ws

Unix CLI workspace manager. Accessible via the `ws` binary

_Windows isn't supported mainly because I don't use _Command Prompt_, _PowerShell_, etc., PRs for
that are welcome._

This tool outputs the path of a given workspace so your shell can use to change directories. This is
necessary because processes change working directory only in the context in which they run, the
directory changes back to what it was before the program ran when it exits.

## Installation

Run on your terminal `go get github.com/andradei/ws`

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

When using `ws [workspace name]`, `ws` will output the directory of for the given workspace name.
So you can do the following:

```
cd $(ws myworkspace)
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
