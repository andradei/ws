package main

import (
	"fmt"
	"os"
	"testing"

	color "github.com/fatih/color"
)

func TestGetMetadata(t *testing.T) {
	var metadata = []struct {
		name string
		path string
	}{
		{"code", "/some/path/to/code"},
		{"exercises", "/some/path/to/exercises"},
		{"github", "/some/path/to/github"},
		{"gitlab", "/some/path/to/gitlab"},
		{"mysite", "/some/path/to/mysite"},
	}

	// The metadata file, located at testdata/ws.json must have content that matches the
	// metadata table above.
	var m, err = getMetadata("testdata")
	if err != nil {
		t.Fatalf("want: nil' | got: %v", err)
	}

	t.Run("Path", func(t *testing.T) {
		if m.path != "testdata/ws.json" {
			t.Error("wrong path: want testdata/ws.json | got", m.path)
		}
	})

	t.Run("Content", func(t *testing.T) {
		for i, w := range m.workspaces {
			if w.Name != metadata[i].name {
				t.Errorf("wrong workspace name: want %v | got %v", w.Name, metadata[i].name)
			}

			if w.Path != metadata[i].path {
				t.Errorf("wrong workspace path: want %v | got %v", w.Path, metadata[i].path)
			}
		}
	})

	t.Run("List", func(t *testing.T) {
		var l, err = m.list()
		if err != nil {
			t.Error("(*metadata).list() failed:", err)
		}

		yellow := color.New(color.FgYellow).SprintfFunc()
		green := color.New(color.FgGreen).SprintfFunc()

		// The string formatting is used because the terminal colors need to be compared
		// because they are text that get escaped by the terminal.
		var list = fmt.Sprintf(`  %v
    %v
  %v
    %v
  %v
    %v
  %v
    %v
  %v
    %v
`,
			green("code"),
			yellow("/some/path/to/code"),
			green("exercises"),
			yellow("/some/path/to/exercises"),
			green("github"),
			yellow("/some/path/to/github"),
			green("gitlab"),
			yellow("/some/path/to/gitlab"),
			green("mysite"),
			yellow("/some/path/to/mysite"),
		)

		if l != list {
			t.Errorf("wrong string for list of workspaces, want:\n%v\ngot:\n%v\nIn bytes: want:\n%v\ngot:\n%v\n",
				list, l, []byte(list), []byte(l))
		}
	})
}

func TestGetMetadataCreateFile(t *testing.T) {
	const testDir = "test_generated_directory"
	var _, err = getMetadata(testDir)
	if err != nil {
		t.Error("failed to generate metadata file:", err)
	}

	// Delete test_generated directory
	err = os.RemoveAll(testDir)
	if err != nil {
		t.Fatalf("failed to delete %v directory: %v", testDir, err)
	}
}

func TestInsertAndDelete(t *testing.T) {
	var md, err = getMetadata("testdata")
	if err != nil {
		t.Error("getMetadata failed:", err)
	}

	err = md.insert("testWorkspace", "/test/workspace")
	if err != nil {
		t.Error("insertion failure:", err)
	}

	err = md.delete("testWorkspace")
	if err != nil {
		t.Error("deletion failure:", err)
	}
}
