package scheme

import (
	"testing"
)

func init() {
	Init()
	AddToLoadPath("../")
}

func TestVersion(t *testing.T) {
	var (
		expect = "3.0.7"
		got    = Version().ToString()
	)
	if expect != got {
		t.Errorf("Expect %s got %s", expect, got)
	}
}

func TestList(t *testing.T) {
	var (
		got, err = EvalString("%load-path")
	)
	if err != nil {
		t.Fatal(err)
	}
	if !got.IsList() {
		t.Errorf("Expext true got %t", got.IsList())
	}
}

func TestSocketPath(t *testing.T) {
	UseModule("go server")
	var (
		got, err = Eval("socket-file")
		expect   = "/tmp/go-scheme.socket"
	)
	if err != nil {
		t.Fatal(err)
	}
	if got.String() != expect {
		t.Errorf("Expect %s got %s", expect, got)
	}
}

func TestEval(t *testing.T) {
	var (
		got, err = Eval("(version)")
		expect   = "3.0.7"
	)
	if err != nil {
		t.Fatal(err)
	}
	if got.ToString() != expect {
		t.Errorf("Expect %s got %s", got.ToString(), expect)
	}
}

func TestEvalFail(t *testing.T) {
	var (
		_, err = Eval("(versio)")
	)
	if err == nil {
		t.Fatalf("Expected error got nil")
	}
}
