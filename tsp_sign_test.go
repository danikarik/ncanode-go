package ncanode_test

import (
	"testing"

	"github.com/danikarik/ncanode-go"
	"github.com/stretchr/testify/require"
)

func TestTSPSign(t *testing.T) {
	client, err := ncanode.NewClient("http://127.0.0.1:14579")
	require.NoError(t, err)

	testCases := []struct {
		Name   string
		Raw    string
		Policy ncanode.Policy
		Algo   ncanode.HashAlgorithm
	}{
		{
			Name: "Default",
			Raw:  _defaultRaw,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			resp, err := client.TSPSign(tc.Raw, tc.Policy, tc.Algo)
			require.NoError(t, err)
			require.NotEmpty(t, resp.Result.TSP)
		})
	}
}
