package ncanode_test

import (
	"testing"

	"github.com/danikarik/ncanode-go/v2"
	"github.com/stretchr/testify/require"
)

func TestPKCS12Info(t *testing.T) {
	client, err := ncanode.NewClient("http://127.0.0.1:14579")
	require.NoError(t, err)

	testCases := []struct {
		Name           string
		Path           string
		CheckOCSP      bool
		CheckCRL       bool
		Alias          string
		ExpectedResult bool
	}{
		{
			Name:           "Personal/Active/Auth",
			Path:           "personal/active/AUTH_RSA256_b2c70a9ff7a5dc59de0ccc3c0ddde3437cc1f12d.p12",
			Alias:          "b2c70a9ff7a5dc59de0ccc3c0ddde3437cc1f12d",
			ExpectedResult: true,
		},
		{
			Name:           "Personal/Active/Sign",
			Path:           "personal/active/RSA256_11c707bd54cbfcccb3815e39c4eb57b1dc7dfea9.p12",
			Alias:          "11c707bd54cbfcccb3815e39c4eb57b1dc7dfea9",
			ExpectedResult: true,
		},
		{
			Name:           "Personal/Revoked/Auth",
			Path:           "personal/revoked/AUTH_RSA256_60d5a0346dd52fb5f3b9148e6cfbcd6cf323d119.p12",
			CheckOCSP:      true,
			CheckCRL:       true,
			Alias:          "60d5a0346dd52fb5f3b9148e6cfbcd6cf323d119",
			ExpectedResult: false,
		},
		{
			Name:           "Personal/Revoked/Sign",
			Path:           "personal/revoked/RSA256_346a7e57c2995259140b6fc375b6ff3bba7e852f.p12",
			CheckOCSP:      true,
			CheckCRL:       true,
			Alias:          "346a7e57c2995259140b6fc375b6ff3bba7e852f",
			ExpectedResult: false,
		},
		{
			Name:           "Organization/Active/Auth",
			Path:           "organization/active/AUTH_RSA256_d682726b55b3e62600ea8fe1c74e75fc96f47790.p12",
			Alias:          "d682726b55b3e62600ea8fe1c74e75fc96f47790",
			ExpectedResult: true,
		},
		{
			Name:           "Organization/Active/Sign",
			Path:           "organization/active/GOSTKNCA_13a3741b52d7bf860aab199b56fad1f4652e357b.p12",
			Alias:          "13a3741b52d7bf860aab199b56fad1f4652e357b",
			ExpectedResult: true,
		},
		{
			Name:           "Organization/Revoked/Auth",
			Path:           "organization/revoked/AUTH_RSA256_7d6d313ac5bf7367a2a69f28607a8deb80dd3ba9.p12",
			CheckOCSP:      true,
			CheckCRL:       true,
			Alias:          "7d6d313ac5bf7367a2a69f28607a8deb80dd3ba9",
			ExpectedResult: false,
		},
		{
			Name:           "Organization/Revoked/Sign",
			Path:           "organization/revoked/GOSTKNCA_c372b7809440fcb681051bbc89ea2f089fd2fe16.p12",
			CheckOCSP:      true,
			CheckCRL:       true,
			Alias:          "c372b7809440fcb681051bbc89ea2f089fd2fe16",
			ExpectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			p12, err := base64content(tc.Path)
			require.NoError(t, err)

			resp, err := client.PKCS12Info(p12, _defaultPassword, tc.CheckOCSP, tc.CheckCRL, tc.Alias)
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
