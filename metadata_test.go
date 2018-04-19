package main

import (
	"testing"
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

	var _, err1 = getMetadata("probablyInvalidDirectoryUsedForTesting", false)
	if err1 == nil {
		t.Fatalf("want: unable to read metadata: ...' | got: nil | Actual error: %v", err1)
	}

	m, err2 := getMetadata("testdata", false)
	if err2 != nil {
		t.Fatalf("want: nil' | got: %v", err2)
	}

	if m.path != "testdata/ws.json" {
		t.Error("wrong path: want testdata/ws.json | got", m.path)
	}

	for i, w := range m.workspaces {
		if w.Name != metadata[i].name {
			t.Errorf("wrong workspace name: want %v | got %v", w.Name, metadata[i].name)
		}

		if w.Path != metadata[i].path {
			t.Errorf("wrong workspace path: want %v | got %v", w.Path, metadata[i].path)
		}
	}
}

func TestInsertAndDelete(t *testing.T) {
	var md, err = getMetadata("testdata", false)

	err = md.insert("testWorkspace", "/test/workspace")
	if err != nil {
		t.Error("insertion failure", err)
	}

	err = md.delete("testWorkspace")
	if err != nil {
		t.Error("deletion failure", err)
	}
}
