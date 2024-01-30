//go:build go1.18

package optional_test

import (
	"fluentqa-go-tutorials/qfluentlib/optional"
	"testing"

	"github.com/gookit/goutil/dump"
)

func TestOptFactory_of(t *testing.T) {
	opt := optional.Of(0)

	var a int
	a = opt.OrElseGet(34)

	dump.P(a, opt.OrElseGet(34))
}
