package ncanode_test

import (
	"errors"
	"testing"

	"github.com/danikarik/ncanode-go"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	_, err := ncanode.NewClient("http://127.0.0.1:8080")
	require.Error(t, err)
	require.True(t, errors.Is(err, ncanode.ErrFailedConnection))
}
