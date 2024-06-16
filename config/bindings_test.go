package config

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.ngrok.com/ngrok/internal/tunnel/proto"
)

func testBindings[T tunnelConfigPrivate, O any, OT any](
	t *testing.T,
	makeOpts func(...OT) Tunnel,
	getBindings func(*O) []string,
) {
	optsFunc := func(opts ...any) Tunnel {
		return makeOpts(assertSlice[OT](opts)...)
	}

	cases := testCases[T, O]{
		{
			name: "absent",
			opts: optsFunc(),
			expectOpts: func(t *testing.T, opts *O) {
				actual := getBindings(opts)
				require.Empty(t, actual)
			},
		},
		{
			name: "with bindings",
			opts: optsFunc(WithBindings("public")),
			expectOpts: func(t *testing.T, opts *O) {
				actual := getBindings(opts)
				// require.NotEmpty(t, actual)
				require.Equal(t, []string{"public"}, actual)
			},
		},
	}

	cases.runAll(t)
}

func TestBindings(t *testing.T) {
	testBindings[*httpOptions](t, HTTPEndpoint, func(opts *proto.HTTPEndpoint) []string {
		return opts.Bindings
	})
	testBindings[*tlsOptions](t, TLSEndpoint, func(opts *proto.TLSEndpoint) []string {
		return opts.Bindings
	})
	testBindings[*tcpOptions](t, TCPEndpoint, func(opts *proto.TCPEndpoint) []string {
		return opts.Bindings
	})
}
