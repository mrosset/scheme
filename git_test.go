package scheme

import (
	"testing"

	"github.com/mrosset/util/file"
	"os"
)

func TestGitVersion(t *testing.T) {
	t.Parallel()
	var (
		git    = NewGitCommand()
		expect = "git version 2.34.1"
		got    = git.Version()
	)
	if got != expect {
		t.Errorf("Expect %v but got %v", expect, got)
	}
}

func TestClone(t *testing.T) {
	t.Parallel()
	var (
		git    = NewGitCommand()
		expect = "testdata/git/scheme"
	)
	defer os.RemoveAll(expect)
	err := git.Clone("./", expect)
	if err != nil {
		t.Fatal(err)
	}
	if !file.Exists(expect) {
		t.Errorf("Expect directory %s to exist", expect)
	}
}
