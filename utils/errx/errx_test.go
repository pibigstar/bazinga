package errx_test

import (
	"github.com/pibigstar/bazinga/internal/code"
	"github.com/pibigstar/bazinga/utils/errx"
	"testing"
)

func TestErrX(t *testing.T) {
	err := errx.NewWithCode(code.Error_Internal)
	t.Log(err)

	err = errx.NewWithMsg("Hello: %s", "world")
	t.Log(err)
}
