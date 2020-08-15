package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpackString(t *testing.T) {
	require.Equal(t, "aaaabccddddde", unpackString("a4bc2d5e"))
	require.Equal(t, "abcd", unpackString("abcd"))
	require.Equal(t, "", unpackString("45"))
}
