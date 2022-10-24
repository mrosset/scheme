package scheme

import (
	"os"
	"path/filepath"
)

var (
	prefix = filepath.Join(os.Getenv("HOME"), ".local/scheme")
)

func Prefix() string {
	return prefix
}

func SetPrefix(path string) {
	prefix = path
}

func UseGit(uri, module string) error {
	var (
		git  = NewGitCommand()
		path = filepath.Join(prefix, filepath.Base(uri))
	)
	if err := git.Clone(uri, path); err != nil {
		return err
	}
	UseModule(module)
	return nil
}

func UseGitBranch(uri, branch string) {
}
