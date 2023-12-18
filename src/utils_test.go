package src_test

import (
	"testing"

	"github.com/psycomentis/psycofolio++/src"
)

func TestGetFullPath(t *testing.T) {
	t.Log(src.GetFullPath("$HOME/first/last"))
	t.Log(src.GetFullPath("$green/first/last"))
	t.Log(src.GetFullPath("$CPU/first/last/$USR"))
}
