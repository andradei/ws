package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	color "github.com/fatih/color"
)

const file = "ws.json"

type workspace struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type metadata struct {
	workspaces []workspace
	// Metadata filepath
	path string
}

// Try to read the contents on the metadata file and decode it into the workspaces type.
// Create the metadata file if it didn't exist previously. This function may not write to the
// metadata file.
func getMetadata(path string) (*metadata, error) {
	path = filepath.Join(path, file)

	md, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, 0700)
			if err != nil {
				return nil, fmt.Errorf("unable to create metadata config directory %v: %v", path, err)
			}
			// metadata will have an empty workspaces instance since no metadata file exists yet.
			return &metadata{make([]workspace, 0, 1), path}, nil
		}
		return nil, fmt.Errorf("unable to open metadata file: %v", err)
	}
	defer func() {
		if e := md.Close(); e != nil {
			panic(fmt.Sprintf("init: failed to close metadata file: %v", e))
		}
	}()

	ws, err := parseWorkspaces(path)
	if err != nil {
		return nil, err
	}

	return &metadata{ws, path}, nil
}

// Parse a file looking for compatible metadata. Returns a slice of workspaces if able, the slice
// may be empty if the file is empty or contains only the file metadata information.
func parseWorkspaces(path string) ([]workspace, error) {
	// Parse JSON metadata.
	var ws []workspace
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read metadata: %v", err)
	}
	err = json.Unmarshal(b, &ws)
	if err != nil {
		return nil, fmt.Errorf("unable to decode metadata: %v", err)
	}

	return ws, nil
}

func (md *metadata) insert(name string, path string) error {
	if _, err := md.getWorkspace(name); err == nil {
		return fmt.Errorf("workspace name already exists: %v", err)
	}

	md.workspaces = append(md.workspaces, workspace{name, path})
	if err := md.save(); err != nil {
		return fmt.Errorf("unable to insert workspace: %v", err)
	}

	return nil
}

func (md *metadata) save() error {
	b, err := json.MarshalIndent(md.workspaces, "", "  ")
	if err != nil {
		return fmt.Errorf("save: unable to encode workspace: %v", err)
	}

	err = ioutil.WriteFile(md.path, b, 0600)
	if err != nil {
		return fmt.Errorf("save: unable to write metadata: %v", err)
	}

	return nil
}

func (md *metadata) getWorkspace(name string) (int, error) {
	for i, w := range md.workspaces {
		if w.Name == name {
			return i, nil
		}
	}

	return 0, fmt.Errorf("workspace %s doesn't exist", name)
}

func (md *metadata) delete(name string) error {
	i, err := md.getWorkspace(name)
	if err != nil {
		return fmt.Errorf("delete: unable to retrieve workspace: %v", err)
	}

	md.workspaces = append(md.workspaces[:i], md.workspaces[i+1:]...)
	err = md.save()
	if err != nil {
		return fmt.Errorf("unable to delete workspace: %v", err)
	}
	return nil
}

// Create a list of workspaces stored in the metadata file.
func (md *metadata) list() (string, error) {
	var result bytes.Buffer
	yellow := color.New(color.FgYellow).SprintfFunc()
	green := color.New(color.FgGreen).SprintfFunc()

	for _, ws := range md.workspaces {
		_, err := result.WriteString(fmt.Sprintf("  %s\n    %s\n", green(ws.Name), yellow(ws.Path)))
		if err != nil {
			return "", fmt.Errorf("unable to write list: %v", err)
		}
	}

	return result.String(), nil
}
