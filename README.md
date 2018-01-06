# ws

Unix CLI directory alias manager. Accessible via the `ws` binary

_Windows isn't supported mainly because I don't use _Command Prompt_, _PowerShell_, etc., PRs for
that are welcome._

In `ws` directories are called workspaces.

## Installation

Run on your terminal `go get github.com/andradei/ws`. Or download a version from the Releases page.

## Commands

TODO: Create a gif with usage examples

```bash
$ ws -help

ws - Directory alias manager

Usage: ws <command [workspace name]> | <workspace name>

  -create | -c <workspace name>
       Create a workspace with given name

   -delete | -d <workspace name>
       Delete an existing workspace by name

   -help | -h
       Display this help message

   -list | -l
       List existing workspaces

   -version | -v
       Display program version

```

When using the `ws <workspace name>` form, `ws` will output the directory of for the given workspace name.
So you can do the following:

```
cd $(ws my_workspace)
```

## Examples

- Create workspace: `ws -create project1`

- Go to workspace: `cd $(ws project1)`

- Delete workspace: `ws -delete project1`
