package ncanode_test

import (
	"testing"

	"github.com/danikarik/ncanode-go"
	"github.com/stretchr/testify/require"
)

func TestXMLSign(t *testing.T) {
	client, err := ncanode.NewClient("http://127.0.0.1:14579")
	require.NoError(t, err)

	testCases := []struct {
		Name   string
		Path   string
		Config *ncanode.TSPConfig
	}{
		{
			Name: "Personal/Active/Auth",
			Path: "personal/active/AUTH_RSA256_b2c70a9ff7a5dc59de0ccc3c0ddde3437cc1f12d.p12",
		},
		{
			Name: "Personal/Active/Sign",
			Path: "personal/active/RSA256_11c707bd54cbfcccb3815e39c4eb57b1dc7dfea9.p12",
		},
		{
			Name: "Organization/Active/Auth",
			Path: "organization/active/AUTH_RSA256_d682726b55b3e62600ea8fe1c74e75fc96f47790.p12",
		},
		{
			Name: "Organization/Active/Sign",
			Path: "organization/active/GOSTKNCA_13a3741b52d7bf860aab199b56fad1f4652e357b.p12",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			p12, err := base64content(tc.Path)
			require.NoError(t, err)

			resp, err := client.XMLSign(p12, _defaultPassword, _defaultXML, tc.Config)
			require.NoError(t, err)
			require.NotEmpty(t, resp.Result.XML)

			if tc.Config != nil && tc.Config.Enabled {
				require.NotEmpty(t, resp.Result.TSP)
			}
		})
	}
}
