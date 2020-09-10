package ncanode_test

import (
	"testing"

	"github.com/danikarik/ncanode-go/v2"
	"github.com/stretchr/testify/require"
)

func TestX509Info(t *testing.T) {
	client, err := ncanode.NewClient("http://127.0.0.1:14579")
	require.NoError(t, err)

	testCases := []struct {
		Name           string
		Path           string
		CheckOCSP      bool
		CheckCRL       bool
		ExpectedResult bool
	}{
		{
			Name:           "Personal/Active/Auth",
			Path:           "personal/active/AUTH_RSA256_b2c70a9ff7a5dc59de0ccc3c0ddde3437cc1f12d.cer",
			ExpectedResult: true,
		},
		{
			Name:           "Personal/Active/Sign",
			Path:           "personal/active/RSA256_11c707bd54cbfcccb3815e39c4eb57b1dc7dfea9.cer",
			ExpectedResult: true,
		},
		{
			Name:           "Organization/Active/Auth",
			Path:           "organization/active/AUTH_RSA256_d682726b55b3e62600ea8fe1c74e75fc96f47790.cer",
			ExpectedResult: true,
		},
		{
			Name:           "Organization/Active/Sign",
			Path:           "organization/active/GOSTKNCA_13a3741b52d7bf860aab199b56fad1f4652e357b.cer",
			ExpectedResult: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			cert, err := base64content(tc.Path)
			require.NoError(t, err)

			resp, err := client.X509Info(cert, tc.CheckOCSP, tc.CheckCRL)
			require.NoError(t, err)
			require.Equal(t, tc.ExpectedResult, resp.Cert.Valid)

			if tc.CheckOCSP {
				require.NotNil(t, resp.Cert.OCSP)
			}

			if tc.CheckCRL {
				require.NotNil(t, resp.Cert.CRL)
			}
		})
	}
}
