package scheme

import (
	"testing"
)

const (
	GUILE_VERSION = "3.0.7"
)

func init() {
	Init()
	AddToLoadPath("../")
}

func TestGuileVersion(t *testing.T) {
	var (
		expect = GUILE_VERSION
		got    = GuileVersion().String()
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
		t.Error("Expect list type")
	}
}

func TestBool(t *testing.T) {
	var (
		got, err = EvalString("#t")
	)
	if err != nil {
		t.Fatal(err)
	}
	if !got.IsBool() {
		t.Error("Expect boolean type")
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
		expect   = GUILE_VERSION
	)
	if err != nil {
		t.Fatal(err)
	}
	if got.String() != expect {
		t.Errorf("Expect %s got %s", got.ToString(), expect)
	}
}

func TestEvalFail(t *testing.T) {
	var (
		_, err = Eval("(versio)")
	)
	if err == nil {
		t.Errorf("Expected error got nil")
	}
}
