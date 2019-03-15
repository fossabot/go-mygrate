package generate

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

func TestInit(t *testing.T) {
	tmpDir := os.TempDir()
	path := tmpDir + "/_TestInit_/"
	err := Init(path)
	if err != nil {
		t.Fatalf("Init %q: %s", path, err)
	}
	defer os.RemoveAll(path)

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("os.Stat %q: %s", path, err)
	}

	if !info.IsDir() {
		t.Fatalf("Init %q: did not create directory", path)
	}
}

func TestRenderMygration(t *testing.T) {
	const id = 1
	const name = "init"

	pathGoldenFile := filepath.Join("testdata", t.Name()+".golden")

	if *update {
		t.Log("renderMygration: update golden file")
		buf, err := renderMygration(id, name)
		if err != nil {
			t.Fatalf("renderMygration: failed to update golden file: %s", err)
		}

		if err := ioutil.WriteFile(pathGoldenFile, buf.Bytes(), 0644); err != nil {
			t.Fatalf("renderMygration: failed to update golden file: %s", err)
		}
	}

	buf, err := renderMygration(id, name)
	if err != nil {
		t.Fatalf("renderMygration: err %s", err)
	}

	generatedFile, err := ioutil.ReadFile(pathGoldenFile)
	if err != nil {
		t.Fatalf("renderMygration: err %s", err)
	}

	if !bytes.Equal(buf.Bytes(), []byte(generatedFile)) {
		t.Fatal("renderMygration: mygration does not match golden file")
	}
}

func TestGenerateMygration_File(t *testing.T) {
	tmpDir := os.TempDir()
	path := tmpDir + "/_TestGenerateMygration_File_/"
	err := os.MkdirAll(path, 0644)
	if err != nil {
		t.Fatalf("GenerateMygration %q: err %s", path, err)
	}
	defer os.RemoveAll(path)

	err = GenerateMygration(path, 1, "init")
	if err != nil {
		t.Fatalf("GenerateMygration %q: err %s", path, err)
	}

	_, err = os.Stat(path + "/1_init.go")
	if os.IsNotExist(err) {
		t.Fatalf("GenerateMygration %q: did not create migration file", path)
	}

	if err != nil {
		t.Fatalf("GenerateMygration %q: err %s", path, err)
	}
}
