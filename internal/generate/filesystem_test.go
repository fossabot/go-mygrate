package generate

import (
	"os"
	"testing"
)

func TestEnsureDir_CreatesDirectory(t *testing.T) {
	tmpDir := os.TempDir()
	path := tmpDir + "/_TestEnsureDir_CreatesDirectory_/"
	err := ensureDir(path)
	if err != nil {
		t.Fatalf("ensureDir %q: %s", path, err)
	}
	defer os.RemoveAll(path)

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("os.Stat %q: %s", path, err)
	}

	if !info.IsDir() {
		t.Fatalf("ensureDir %q: did not create directory", path)
	}
}

func TestFindMigrationsInDir(t *testing.T) {
	migrations := findMygrationsInDir("./testdata/TestFindMigrationsInDir")
	if len(migrations) != 2 {
		t.Fatalf("findMigrationsInDir: did not find migrations")
	}
}

func TestFindMigrationsInDir_Pattern(t *testing.T) {
	migrations := findMygrationsInDir("./testdata/TestFindMigrationsInDir_Pattern")
	if len(migrations) != 2 {
		t.Fatalf("findMigrationsInDir: did not find migrations")
	}
}

func TestParseIDAndName(t *testing.T) {
	path := "/example/1_init.go"

	id, name, err := parseIDAndName(path)
	if err != nil {
		t.Fatalf("parseIDAndName %q: err %s", path, err)
	}

	if id != 1 {
		t.Fatalf("parseIDAndName %q: id is not 1", path)
	}

	if name != "init" {
		t.Fatalf("parseIDAndName %q: name is not 'init'", path)
	}
}

func TestParseIDAndName_WrongFormat(t *testing.T) {
	path := "/example/init.go"

	_, _, err := parseIDAndName(path)
	if err == nil {
		t.Fatalf("parseIDAndName %q: excpected error", path)
	}
}

func TestParseIDAndName_WrongID(t *testing.T) {
	path := "/example/a_init.go"

	_, _, err := parseIDAndName(path)
	if err == nil {
		t.Fatalf("parseIDAndName %q: excpected error", path)
	}
}
