package scheme

import (
	. "github.com/mrosset/test"
	"testing"
)

func TestVersion(t *testing.T) {
	Test{
		Expect: "2.2.4",
		Got:    Version().ToString(),
	}.Equals(t)
}

func TestAddToLoadPath(t *testing.T) {
	// Test{
	//	Expect: true,
	//	Got:    AddToLoadPath("scm").Bool(),
	// }.Equals(t)
}
