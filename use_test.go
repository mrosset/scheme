package scheme

import (
	"github.com/mrosset/util/file"
	"os"
	"path/filepath"
	"testing"
)

func TestPrefix(t *testing.T) {
	if Prefix() != filepath.Join(os.Getenv("HOME"), ".local/scheme") {
		t.Errorf("expect $HOME/.local/scheme got %s", Prefix())
	}
}

func TestUse(t *testing.T) {
	var (
		expect = "testdata"
	)
	SetPrefix(expect)
	if file.Exists(expect) {
		os.RemoveAll(expect)
	}
	defer os.RemoveAll(expect)
	if Prefix() != expect {
		t.Fatalf("Expect prefix %s got %s", expect, Prefix())
	}
	err := UseGit("./", "test use")
	if err != nil {
		t.Fatal(err)
	}
	res, err := EvalString("use-variable")
	if err != nil {
		t.Fatal(err)
	}
	if !res.Bool() {
		t.Error("expected use-variable to be true")
	}
}
