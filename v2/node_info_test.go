package ncanode_test

import (
	"testing"

	"github.com/danikarik/ncanode-go/v2"
	"github.com/stretchr/testify/require"
)

func TestNodeInfo(t *testing.T) {
	client, err := ncanode.NewClient("http://127.0.0.1:14579")
	require.NoError(t, err)

	resp, err := client.NodeInfo()
	require.NoError(t, err)

	require.False(t, resp.Result.Datetime.IsZero())
	require.NotEmpty(t, resp.Result.Timezone)
	require.Contains(t, resp.Result.Name, "NCANode")
	require.NotEmpty(t, resp.Result.Version)
}
