package scheme

import (
	"os/exec"
	"strings"
)

// Provides Cmd type for git program
type GitCommand exec.Cmd

// // Returns a new GitCommand
func NewGitCommand() GitCommand {
	var (
		bin, err = exec.LookPath("git")
	)
	if err != nil {
		panic(err)
	}
	return GitCommand{Path: bin}
}

// Clones git URI to DIR path
func (g GitCommand) Clone(uri, dir string) error {
	var (
		cmd = exec.Cmd(g)
	)
	cmd.Args = []string{cmd.Path, "clone", uri, dir}
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// Returns the git command version.
func (g GitCommand) Version() string {
	var (
		cmd = exec.Cmd(g)
	)
	cmd.Args = []string{cmd.Path, "version"}
	out, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return strings.TrimSpace(string(out[:]))
}
