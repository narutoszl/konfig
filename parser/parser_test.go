package parser

import (
	io "io"
	"testing"

	"github.com/lalamove/konfig"
	"github.com/stretchr/testify/require"
)

func TestParserFunc(t *testing.T) {
	var ran bool
	var f = Func(func(r io.Reader, s konfig.Values) error {
		ran = true
		return nil
	})
	f.Parse(nil, nil)
	require.Equal(t, true, ran)
}
