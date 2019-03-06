package scheme

import (
	. "github.com/mrosset/via/pkg/test"
	"testing"
)

func TestVersion(t *testing.T) {
	Test{
		Expect: "2.2.4",
		Got:    Version().String(),
	}.Equals(t)
}
